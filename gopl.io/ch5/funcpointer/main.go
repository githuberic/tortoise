package main

import "fmt"

func square(n int) int {
	return n * n
}

func main()  {
	// 函数变量
	f := square
	fmt.Println(f(3))

	var fun func(int) int
	if fun == nil {
		fun = square
	}
	fmt.Println(fun(4))
}
