package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var signs = make(chan os.Signal, 1)
	var done = make(chan bool, 1)

	// `signal.Notify`在给定的channel上面注册该channel
	// 可以接受的信号
	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM)

	// 这个goroutine阻塞等待信号的到来，当信号到来的时候，
	// 输出该信号，然后通知程序可以结束了
	go func() {
		sig := <-signs
		fmt.Println(sig)
		done <- true
	}()

	// 程序将等待接受信号，然后退出
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
