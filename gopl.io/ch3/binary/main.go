package main

import "fmt"

/**
to do...
 */
func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	// %b 以二进制形式输出，08：输出结果补0，补够8位
	fmt.Printf("%08b,%d\n", x, x)
	fmt.Printf("%08b,%d\n", y, y)
	fmt.Printf("%08b,%d\n", x&y, x&y)
	fmt.Printf("%08b,%d\n", x|y, x|y)
	// 按位异或的3个特点:
	// (1) 0^0=0,0^1=1  0异或任何数＝任何数
	// (2) 1^0=1,1^1=0  1异或任何数=任何数取反
	fmt.Printf("%08b,%d\n", x^y, x^y) //^按位异或(XOR)，若作为1元操作符,它标识按位取反或按位取补,0和1异或0都不变，异或1则取反

	fmt.Printf("%08b,%d\n", x&^y, x&^y)
}
