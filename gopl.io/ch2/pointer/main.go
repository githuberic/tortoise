package main

import "fmt"

func main() {
	x := 1

	// &x(x的地址)获取一个指向整形变量的指针,它的类型是整形指针(*int)
	var p *int
	p = &x
	fmt.Println("Address=", &p)
	fmt.Println("Value=", *p)

	*p = 2
	fmt.Println("Value(x)=",x)

	var xx, y int
	fmt.Println(&xx == &xx, &xx == &y, &xx == nil)

	v := 1
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	*p++
	return *p
}
