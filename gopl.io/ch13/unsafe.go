package main

import (
	"fmt"
	"unsafe"
)

var x struct{
	a bool
	b int16
	c []int
}

func main()  {
	fmt.Println(unsafe.Sizeof(float64(0)))
	fmt.Println("=======Sizeof=======")
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Sizeof(x.a))
	fmt.Println(unsafe.Sizeof(x.b))
	fmt.Println(unsafe.Sizeof(x.c))
	fmt.Println("=======Alignof=======")
	fmt.Println(unsafe.Alignof(x))
	fmt.Println(unsafe.Alignof(x.a))
	fmt.Println(unsafe.Alignof(x.b))
	fmt.Println(unsafe.Alignof(x.c))
	fmt.Println("=======Offsetof=======")
	fmt.Println(unsafe.Offsetof(x.a))
	fmt.Println(unsafe.Offsetof(x.b))
	fmt.Println(unsafe.Offsetof(x.c))
}
