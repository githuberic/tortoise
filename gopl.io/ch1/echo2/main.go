package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// 每次迭代 range都会产生一对值 索引/这个索引处元素的值，我们不需要索引时，通过空标识符 _ 来解决
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Print(s)
}
