package bad

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestE1(t *testing.T) {
	var wg sync.WaitGroup
	go dosomething(100, &wg) // 启动第一个goroutine
	go dosomething(110, &wg) // 启动第二个goroutine
	go dosomething(120, &wg) // 启动第三个goroutine
	go dosomething(130, &wg) // 启动第四个goroutine

	wg.Wait() // 主goroutine等待完成
	fmt.Println("Done")
}

func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
	duration := millisecs * time.Millisecond
	time.Sleep(duration) // 故意sleep一段时间

	wg.Add(1)
	fmt.Println("后台执行, duration:", duration)
	wg.Done()
}
