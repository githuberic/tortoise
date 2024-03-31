package incr

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Counter struct {
	count  uint64
	locker sync.RWMutex
}

func (c *Counter) Incr() {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.count++
}

func (c *Counter) Count() uint64 {
	c.locker.RLock()
	defer c.locker.RUnlock()
	return c.count
}

func TestIncr(t *testing.T) {
	var counter Counter

	for i := 0; i < 10; i++ {
		go func() {
			for {
				fmt.Println("Counter=", counter.Count())
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for {
		counter.Incr()
		time.Sleep(time.Second)
	}
}
