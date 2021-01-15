package main

import "fmt"

func main() {
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	bar := foo[1:4]
	bar[1] = 99
	// foo 和 bar 的内存是共享的，所以，foo 和 bar 对数组内容的修改都会影响到对方
	fmt.Println(foo)
	fmt.Println(bar)

	arrA := make([]int, 32)
	arrB := arrA[1:16]
	// 对 a 做了一个 append()的操作，这个操作会让 a 重新分配内存，这就会导致 a 和 b 不再共享
	arrA = append(arrA, 1)
	arrA[2] = 42
	fmt.Println(arrA)
	fmt.Println(arrB)
}
