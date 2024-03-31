package good

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestE2(t *testing.T) {
	var wg sync.WaitGroup

	dosomethingE2(100, &wg) // 调用方法，把计数值加1，并启动任务goroutine
	dosomethingE2(110, &wg) // 调用方法，把计数值加1，并启动任务goroutine
	dosomethingE2(120, &wg) // 调用方法，把计数值加1，并启动任务goroutine
	dosomethingE2(130, &wg) // 调用方法，把计数值加1，并启动任务goroutine

	wg.Wait() // 主goroutine等待，代码逻辑保证了四次Add(1)都已经执行完了
	fmt.Println("Done")
}

func dosomethingE2(millisecs time.Duration, wg *sync.WaitGroup) {
	wg.Add(1) // 计数值加1，再启动goroutine

	go func() {
		duration := millisecs * time.Millisecond
		time.Sleep(duration)
		fmt.Println("后台执行, duration:", duration)
		wg.Done()
	}()
}

