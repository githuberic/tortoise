package main

import "fmt"

func double(x int) int {
	return x + x
}

func doublev2(x int) (result int) {
	defer func() {fmt.Printf("double(%d) = %d \n",x,result)}()
	return x + x
}

func main()  {
	doublev2(4)
}