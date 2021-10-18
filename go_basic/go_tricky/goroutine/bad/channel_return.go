package go_bad

import (
	"fmt"
	"testing"
)

/**
Goroutine泄露

Go语言是带内存自动回收的特性，因此内存一般不会泄漏。但是Goroutine确存在泄漏的情况，
同时泄漏的Goroutine引用的内存同样无法被回收。

上面的程序中后台Goroutine向管道输入自然数序列，main函数中输出序列。
但是当break跳出for循环的时候，后台Goroutine就处于无法被回收的状态了。
*/
func TestChannelReturn(t *testing.T) {
	ch := func() <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				ch <- i
			}
		}()
		return ch
	}()

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}
}
