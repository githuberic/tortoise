package main

import (
	"fmt"
)

const boilingF = 212.0

/**
 变量:type
 常量:const
 类型:type
 函数:func
 */
func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %g℉ or %g℃\n",f,c)
}