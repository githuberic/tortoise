package mutex

import (
	"fmt"
	"sync"
	"testing"
)

type Counter struct {
	sync.Mutex
	Count uint64
}

func TestVerifyV3(t *testing.T) {
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count)
}
