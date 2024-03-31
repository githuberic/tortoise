package dircection

import (
	"fmt"
	"testing"
)

// 这个ping函数只接收能够发送数据的通道作为参数，试图从这个通道接收数据
// 会导致编译错误，这里只写的定义方式为`chan<- string`表示这个类型为
// 字符串的通道为只写通道
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong函数接收两个通道参数，一个是只读的pings，使用`<-chan string`定义
// 另外一个是只写的pongs，使用`chan<- string`来定义
func pong(pings <-chan string, pongs chan<- string) {
	pongs <- <-pings
}

func TestVerify(t *testing.T) {
	var pings = make(chan string, 1)
	var pongs = make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
