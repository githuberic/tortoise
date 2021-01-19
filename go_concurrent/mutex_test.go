package go_concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var num = 0
func TestVerifyMutex(t *testing.T) {
	mu := sync.Mutex{}
	for i := 0; i < 10000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			num += 1
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("num=",num)
}
