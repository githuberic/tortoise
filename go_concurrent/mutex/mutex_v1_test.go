package mutex

import (
	"fmt"
	"sync"
	"testing"
)

func TestVerifyV1(t *testing.T) {
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var sg sync.WaitGroup
	sg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer sg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	sg.Wait()
	fmt.Println(count)
}
