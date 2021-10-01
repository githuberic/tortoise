package go_interview

import (
	"fmt"
	"testing"
	"time"
)

func TestVerifyCBE1(t *testing.T) {
	var st = time.Now()
	var ch = make(chan bool)

	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	// 无缓冲，发送方阻塞直到接收方接收到数据。
	ch <- true
	fmt.Printf("Cost %0.1f s\n", time.Now().Sub(st).Seconds())
	time.Sleep(time.Second * 5)
}

func TestVerifyCBE2(t *testing.T) {
	var st = time.Now()
	var ch = make(chan bool, 2)

	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	ch <- true
	ch <- true
	// 缓冲区为 2，发送方不阻塞，继续往下执行
	fmt.Printf("Cost %0.1f s\n", time.Now().Sub(st).Seconds()) // cost 0.0 s

	// 缓冲区使用完，发送方阻塞，2s 后接收方接收到数据，释放一个插槽，继续往下执行
	ch <- true
	fmt.Printf("Cost %0.1f s\n", time.Now().Sub(st).Seconds()) // cost 2.0 s
	time.Sleep(time.Second * 5)
}
