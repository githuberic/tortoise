package panic_recover_test

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

// let it crash,往往是应对不确定错误的最好方法
func TestPanicVxExit(t *testing.T) {
	defer func() {
		// 错误恢复，程序没有退出
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
	//fmt.Println("End")
}

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string // 所在函数
}

// 保护方式允许一个函数
func ProtectRun(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()

	entry()
}

func TestPanicRecover(t *testing.T) {
	fmt.Println("运行前")

	ProtectRun(func() {
		fmt.Println("手动宕机前")
		panic(&panicContext{"手动触发panic"})
		fmt.Println("手动宕机后")
	})

	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})

	fmt.Println("运行后")
}
