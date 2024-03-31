package exit

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func doCheckClose(ch chan int) {
	for {
		select {
		case t, closed := <-ch:
			if !closed {
				fmt.Println("taskCh has been closed")
				return
			}
			time.Sleep(time.Millisecond)
			fmt.Printf("task %d is done\n", t)
		}
	}
}
/**
t, beforeClosed := <-taskCh 判断 channel 是否已经关闭，beforeClosed 为 false 表示信道已被关闭。
若关闭，则不再阻塞等待，直接返回，对应的协程随之退出。
 */

func sendTasksCloseCheck()  {
	ch := make(chan int, 10)
	go doCheckClose(ch)
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)
}
func TestDoCheckClose(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	sendTasksCloseCheck()
	time.Sleep(time.Second)
	runtime.GC()
	t.Log(runtime.NumGoroutine())
}