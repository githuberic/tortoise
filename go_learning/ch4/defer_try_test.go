package ch4

import (
	"fmt"
	"testing"
)

func testSub(x int) {
	result := 100 / x
	fmt.Println("result = ", result)
}

func TestDefer(t *testing.T)  {
	defer fmt.Println("a")
	defer fmt.Println("b")

	defer func() {
		testSub(6)
	}()

	defer fmt.Println("c")
}

// 如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行,哪怕函数或某个延迟调用发生错误，这些调用依旧会被执。
// defer 在函数内部使用，一般多用在枷锁、解锁；打开文件、关闭文件等成对出现的情况

