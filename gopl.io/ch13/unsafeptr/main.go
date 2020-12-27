package main

import (
	"fmt"
	"unsafe"
)

// unsafe.Pointer类型是一种特殊类型的指针,它可以存储任何变量的地址
func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func main() {
	//!+main
	var x struct {
		a bool
		b int16
		c []int
	}

	// equivalent to pb := &x.b
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42

	fmt.Println(x.b) // "42"
	//!-main
}
