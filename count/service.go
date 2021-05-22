package count

import (
	"log"

	"sync"

	"proyecto1.com/main/utils"
)

type SafeCounter struct {
	mu sync.Mutex
	V int
}

func (c *SafeCounter) Get() int {
	return c.V
}

func (c *SafeCounter) Restart(n int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	var before = c.V
	c.V = 0
	log.Printf("[Counter] Before: %v. Restarted to 0", before)
	return c.Get()
}

func (c *SafeCounter) Increment(n int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	var before = c.V
	c.V += n
	var after = c.V
	log.Printf("[Counter] Before: %v. After: %v. Incremented by: %v", before, after, n)
	return c.Get()
}

func (c *SafeCounter) Decrement(n int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	var before = c.V
	c.V -= n
	var after = c.V
	log.Printf("[Counter] Before: %v. After: %v. Decremented by: %v", before, after, n)
	return c.Get()
}

func (c *SafeCounter) Print() {
	log.Printf("[Counter] Count: %v", utils.IntToString(c.Get()))
}

var SharedCounter = SafeCounter{V: 0}
