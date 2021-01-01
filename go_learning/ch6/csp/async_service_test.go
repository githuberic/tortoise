package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	//retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func AsyncServiceV2() chan string  {
	 // 下面的这2种声明，执行有区别
	 retChn := make(chan string)
	 //retChn := make(chan string,1)

	 go func() {
	 	ret := service()
	 	fmt.Println("returned result.")
	 	retChn <- ret
		 fmt.Println("service exited.")
	 }()
	 return retChn
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncServiceV2()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}
