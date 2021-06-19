package udpClient

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/adjust/rmq/v3"
)

func InitUDPClientConnection(ip string) (net.Conn, error) {
	address := ip + ":2002"

	resolveAddr, err := net.ResolveUDPAddr("udp4", address)
	c, err := net.DialUDP("udp4", nil, resolveAddr)

	if err != nil {
		fmt.Println(err)

		return nil, errors.New("error connecting to server")
	}

	return c, nil
}

func InvokeUDPClientCall(client net.Conn, operation string, num int) {
	fmt.Fprintf(client, "%v %v\n", operation, num)
}

const (
	prefetchLimit = 1000
	pollDuration  = 100 * time.Millisecond
	numConsumers  = 5

	reportBatchSize = 10000
	consumeDuration = time.Millisecond
	shouldLog       = false
)

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

	count_value := parsed_payload[1]

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
			log.Println("[UDP Server response] La cuenta es de", count_value)
		}
	} else { // reject one per batch
		if err := delivery.Reject(); err != nil {
			log.Printf("failed to reject %s: %s", payload, err)
		} else {
			log.Printf("rejected %s", payload)
		}
	}
}

func ProcessUDPResponses(queueName, ip string) rmq.Queue {
	connection, err := rmq.OpenConnection("consumer", "tcp", ip+":6379", 2, nil)
	if err != nil {
		panic(err)
	}

	queue, err := connection.OpenQueue("responses-" + "UDP Server")
	if err != nil {
		panic(err)
	}

	if err := queue.StartConsuming(prefetchLimit, pollDuration); err != nil {
		panic(err)
	}

	queue.AddConsumer("consumer-udp", NewConsumer())
	return queue
}
