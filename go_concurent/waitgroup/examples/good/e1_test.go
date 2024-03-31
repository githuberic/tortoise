package good

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestE1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(4) // 预先设定WaitGroup的计数值

	go dosomething(100, &wg) // 启动第一个goroutine
	go dosomething(110, &wg) // 启动第二个goroutine
	go dosomething(120, &wg) // 启动第三个goroutine
	go dosomething(130, &wg) // 启动第四个goroutine

	wg.Wait() // 主goroutine等待
	fmt.Println("Done")
}

func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
	duration := millisecs * time.Millisecond
	time.Sleep(duration)

	fmt.Println("后台执行, duration:", duration)
	wg.Done()
}

