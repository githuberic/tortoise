package go_runtime

import (
	"fmt"
	"runtime"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(1) //使用单核
}

func TestGosched1Core(t *testing.T)  {
	//exit := make(chan int)

	go func() {
		//defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i :=0;i < 4;i++ {
		fmt.Println("a:",i)
		if i == 1{
			runtime.Gosched()
		}
	}
	//<-exit
}
