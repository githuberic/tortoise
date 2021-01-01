package empty_interface

import (
	"fmt"
	"testing"
)

// 空指针类型 p interface{}
func DoSomething(p interface{}) {
	// 断言类型 p.(int)
	// if 两段式断言

	// 第一种写法
	/*
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}

	if s, ok := p.(string); ok {
		fmt.Println("stirng", s)
		return
	}
	fmt.Println("Unknow Type")
	*/

	// 第二种写法
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
