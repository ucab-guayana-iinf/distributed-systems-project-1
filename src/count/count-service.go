package count

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	rmq "github.com/adjust/rmq/v3"
	utils "proyecto1.com/main/src/utils"
)

type SafeCounter struct {
	mu sync.Mutex
	V int
}

func (c *SafeCounter) Get() int {
	return c.V
}

func (c *SafeCounter) Restart(source string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	var before = c.V
	c.V = 0
	log.Printf("[Counter] [%s] Before: %v. Restarted to 0", source, before)
	return c.Get()
}

func (c *SafeCounter) Increment(n int, source string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	var before = c.V
	c.V += n
	var after = c.V
	log.Printf("[Counter] [%s] Before: %v. After: %v. Incremented by: %v", source, before, after, n)
	return c.Get()
}

func (c *SafeCounter) Decrement(n int, source string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	var before = c.V
	c.V -= n
	var after = c.V
	log.Printf("[Counter] [%s] Before: %v. After: %v. Decremented by: %v", source, before, after, n)
	return c.Get()
}

func (c *SafeCounter) Print() {
	log.Printf("[Counter] Count: %v", utils.IntToString(c.Get()))
}

var SharedCounter = SafeCounter{V: 0}

const (
	prefetchLimit = 1000
	pollDuration  = 100 * time.Millisecond
	numConsumers  = 5

	reportBatchSize = 10000
	consumeDuration = time.Millisecond
	shouldLog       = false
)

func ProcessMessages() {
	log.Printf("Started processing messages")

	connection, err := rmq.OpenConnection("consumer", "tcp", "localhost:6379", 2, nil)
	if err != nil {
		panic(err)
	}

	queue, err := connection.OpenQueue("operations")
	if err != nil {
		panic(err)
	}

	if err := queue.StartConsuming(prefetchLimit, pollDuration); err != nil {
		panic(err)
	}

	queue.AddConsumer("consumer1", NewConsumer())
}

type Consumer struct {
	name   string
	count  int
	before time.Time
}

func NewConsumer() *Consumer {
	return &Consumer{
		count:  0,
		before: time.Now(),
	}
}

func (consumer *Consumer) Consume(delivery rmq.Delivery) {
	payload := delivery.Payload()
	parsed_payload := strings.Split(payload, ";")

	operation := parsed_payload[0]
	source := parsed_payload[1]
	param := utils.StringToInt(parsed_payload[2])

	consumer.count++
	if consumer.count%reportBatchSize == 0 {
		duration := time.Now().Sub(consumer.before)
		consumer.before = time.Now()
		perSecond := time.Second / (duration / reportBatchSize)
		log.Printf("%s consumed %d %s %d", consumer.name, consumer.count, payload, perSecond)
	}

	if consumer.count%reportBatchSize > 0 {
		if err := delivery.Ack(); err != nil {
			log.Printf("failed to ack %s: %s", payload, err)
		} else {
			ProcessOperation(operation, source, param)
		}
	} else { // reject one per batch
		if err := delivery.Reject(); err != nil {
			log.Printf("failed to reject %s: %s", payload, err)
		} else {
			log.Printf("rejected %s", payload)
		}
	}
}

func ProcessOperation(operation string, source string, param int) {
	switch operation {
	case "Increment":
		SharedCounter.Increment(param, source)
	case "Decrement":
		SharedCounter.Decrement(param, source)
	case "Restart":
		SharedCounter.Restart(source)
	}
}

func Produce(operation string, source string, param int) {
	connection, err := rmq.OpenConnection("producer", "tcp", "localhost:6379", 2, nil)
	if err != nil {
		panic(err)
	}

	queue, err := connection.OpenQueue("operations")
	if err != nil {
		panic(err)
	}

	delivery := fmt.Sprintf("%v;%v;%v", operation, source, param)
	if err := queue.Publish(delivery); err != nil {
		log.Printf("failed to publish: %s", err)
	}
}