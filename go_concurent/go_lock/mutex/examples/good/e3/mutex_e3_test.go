package e3

import (
	"fmt"
	"sync"
	"testing"
)

func TestVerifyV4(t *testing.T) {
	var counterV4 = &CounterV4{}

	var wg = &sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counterV4.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counterV4.Count())
}

type CounterV4 struct {
	mu    sync.Mutex
	count uint64
}

func (c *CounterV4) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *CounterV4) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
