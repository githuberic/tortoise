package main

import "fmt"

/**
defer函数调用的执行顺序与它们分别所属的defer语句的执行顺序相反
 */
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

//!+f
func main() {
	f(3)
}
