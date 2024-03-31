package bad

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
			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}

	// 等待10个goroutine完成
	sg.Wait()

	fmt.Println(count)
}

/**
在这个例子中，我们创建了 10 个 goroutine，同时不断地对一个变量（count）进行加 1 操作，
每个 goroutine 负责执行 10 万次的加 1 操作，
我们期望的最后计数的结果是 10 * 100000 = 1000000 (一百万)。
*/
