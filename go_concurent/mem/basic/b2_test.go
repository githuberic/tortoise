package basic

import "testing"

var ch = make(chan struct{},10)
var s string

func f()  {
	s = "hello world"
	ch <- struct{}{}
}

func TestB2(t *testing.T)  {
	go f()
	<-ch
	println(s)
}

/**
在这个例子中，
s 的初始化（第 5 行）happens before 往 ch 中发送数据，
往 ch 发送数据 happens before 从 ch 中读取出一条数据（第 11 行），
第 12 行打印 s 的值 happens after 第 11 行，所以，打印的结果肯定是初始化后的 s 的值“hello world”。
 */
