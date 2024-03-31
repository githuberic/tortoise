package goroutine

import (
	"fmt"
	"testing"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func TestVerifyE1(t *testing.T) {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 上面的协程在调用之后就异步执行了，所以程序不用等待它们执行完成
	// 就跳到这里来了，下面的Scanln用来从命令行获取一个输入，然后才
	// 让main函数结束
	// 如果没有下面的Scanln语句，程序到这里会直接退出，而上面的协程还
	// 没有来得及执行完，你将无法看到上面两个协程运行的结果
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
