package bad

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func do(ch chan int) {
	for {
		select {
		case t := <-ch:
			time.Sleep(time.Millisecond)
			fmt.Printf("task %d is done\n", t)
		}
	}
}

func sendTasks() {
	ch := make(chan int, 10)
	// sendTasks 中启动了一个子协程 go do(taskCh)
	go do(ch)
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)
}

func TestDo(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	sendTasks()
	time.Sleep(time.Second)
	t.Log(runtime.NumGoroutine())
}
