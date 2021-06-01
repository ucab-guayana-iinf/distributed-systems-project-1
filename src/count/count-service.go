package count

import (
	"log"

	"sync"

	"proyecto1.com/main/src/utils"
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

// TODO: inicializar con el valor que esté en BD
// Si queremos ser exquisitos y sacar 20
var SharedCounter = SafeCounter{V: 0}