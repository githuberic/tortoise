package go_pointer

import (
	"fmt"
	"testing"
)

func double(x int) {
	x += x
}

func TestVerifyE3(t *testing.T) {
	var a = 3
	double(a)
	fmt.Println(a)
}

func doubleV2(x *int) {
	*x += *x
	x = nil
}

func TestVerifyE4(t *testing.T) {
	var a = 3
	doubleV2(&a)
	fmt.Println(a)

	var p = &a
	doubleV2(p)
	fmt.Println(a, p == nil)
}

func TestVerifyV5(t *testing.T) {
	var x struct {
		a int
		b int
		c []int
	}

	var pb = &x.b // 等价于 (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b)
}

// from https://blog.csdn.net/qq_20817327/article/details/117135224