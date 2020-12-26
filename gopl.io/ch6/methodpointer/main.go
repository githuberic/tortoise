package main

import "fmt"

type ExpAdd struct {
	x, y int
}

func (exp ExpAdd) AddExtend(y int) int {
	return exp.x + exp.y + y
}

func main()  {
	var exp = ExpAdd{1,2}
	var add = exp.AddExtend
	fmt.Println(add(3))
}
