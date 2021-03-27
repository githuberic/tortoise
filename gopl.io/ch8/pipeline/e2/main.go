package main

import (
	"fmt"
	"time"
)

func set(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	start := time.Now()

	ch := make(chan int, 20)
	go set(ch)

	for i := range ch {
		fmt.Println("接收到的数据:", i)
	}

	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
