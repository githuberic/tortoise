package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	// %b 以二进制形式输出，08：输出结果补0，补够8位
	fmt.Printf("%08b,%d\n", x, x)
	fmt.Printf("%08b,%d\n", y, y)
	fmt.Printf("%08b,%d\n", x&y, x&y)
	fmt.Printf("%08b,%d\n", x|y, x|y)
	fmt.Printf("%08b,%d\n", x^y, x^y)
	fmt.Printf("%08b,%d\n", x&^y, x&^y)
}
