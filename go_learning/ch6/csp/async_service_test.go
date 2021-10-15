package concurrency

import (
	"fmt"
	"testing"
)

func service() string {
	//time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("Working on something else")
	//time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	// 下面的这2种声明，执行有区别
	retCh := make(chan string)
	//retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	//time.Sleep(time.Second * 1)
}
