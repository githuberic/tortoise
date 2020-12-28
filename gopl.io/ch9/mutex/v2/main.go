package main

import (
	"fmt"
	"sync"
)

var (
	count int
	lock  sync.Mutex
)

func main() {
	for i := 0; i < 2; i++ {
		go func() {
			for i := 1000000; i > 0; i-- {
				lock.Lock()
				count++
				lock.Unlock()
			}
			fmt.Println(count)
		}()
	}

	fmt.Scanf("\n") //等待子线程全部结束
}
