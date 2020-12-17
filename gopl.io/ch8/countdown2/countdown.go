package main

import (
	"fmt"
	"os"
	"time"
)

//!-

func launch() {
	fmt.Println("Lift off!")
}

func main() {
	// ...create abort channel...

	//!-

	//!+abort
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	//!-abort

	//!+
	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	// time.After函数立即返回一个通道，让后启动一个新的goroutine在间隔时间后,发送一个值到它上面
	case <-time.After(10 * time.Second):
		// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}
