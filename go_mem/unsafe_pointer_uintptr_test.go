package go_mem

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestVerifyV5(t *testing.T) {
	foo := &Foo{1,2}
	// 类型
	fmt.Println(foo)
	// 地址
	fmt.Println(&foo)
	// 值
	fmt.Println(*foo)

	base := uintptr(unsafe.Pointer(foo))
	offset := unsafe.Offsetof(foo.A)
	ptr := unsafe.Pointer(base + offset)
	*(*int)(ptr) = 4
	fmt.Println(foo)
}

type Foo struct {
	A int
	B int
}
