package main

import (
	"fmt"
)

func f1() *int {
	t := 3
	return &t
}

func main() {
	x := f1()
	fmt.Println(*x)
}
