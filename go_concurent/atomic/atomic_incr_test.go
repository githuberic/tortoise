package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomicIncr(t *testing.T) {
	var count int32

	var wg = &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&count,1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(count)
}
