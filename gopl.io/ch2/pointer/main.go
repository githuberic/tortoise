package main

import "fmt"

func main()  {
	x := 1
	p := &x
	fmt.Println(x)
	fmt.Println(*p)
	*p =2
	fmt.Println(x)

	var xx, y int
	fmt.Println(&xx == &xx, &xx == &y, &xx == nil)

	v :=1
	fmt.Println(incr(&v))
}

func incr(p *int) int  {
	*p++
	return *p
}
