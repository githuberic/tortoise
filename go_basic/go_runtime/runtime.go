package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("GO Root:", runtime.GOROOT())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Logic CPU nums:", runtime.NumCPU())

	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("最大的可同时使用的CPU核数: ", n)

	go func() {
		fmt.Println("子协程开始执行...")
		//终止当前goroutine
		runtime.Goexit()
		fmt.Println("子协程执行结束...")
	}()

	time.Sleep(1 * time.Second) //主函数休眠3秒，让子协程有充分的时间执行完
	fmt.Println("主函数执行完毕")
}
