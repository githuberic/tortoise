package go_pointer

import (
	"fmt"
	"testing"
	"unsafe"
)

type Demo struct {
	a int32
	b byte
	c int64
}

func TestVerifyV6(t *testing.T)  {
	var p = new(Demo)

	var p1 = (*int32)(unsafe.Pointer(p))
	*p1 = 1

	var p2 = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Sizeof(int32(0))))
	*p2 = 'a'

	var p3 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Sizeof(int64(0))))
	*p3 = 2

	fmt.Println(*p)
}
