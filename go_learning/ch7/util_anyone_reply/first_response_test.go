package concurrency

import (
	"fmt"
	"runtime"
	"testing"
)

func runTask(id int) string {
	//time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10

	// 通过buffer channel，防止go routine泄露
	ch := make(chan string, numOfRunner)

	/**
	运行效果不同
	1:ch := make(chan string)、
	2:ch := make(chan string, numOfRunner)
	*/
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ch <- runTask(i)
		}(i)
	}
	return <-ch
	// (此时 channel 里面 内容没有被读取完毕了，其goroutine没有释放)
}

func TestFirstResponse(t *testing.T) {
	// runtime.NumGoroutine()
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	//time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
