package e1

import (
	"fmt"
	"sync"
	"testing"
)

func TestVerifyV2(t *testing.T) {
	var count = 0

	// 使用WaitGroup等待10个goroutine完成
	var wg = &sync.WaitGroup{}
	wg.Add(10)

	var mutex = &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}
