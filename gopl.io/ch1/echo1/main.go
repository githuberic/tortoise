package main

// os 包提供了一些函数和变量，以与平台无关的方式和操作系统打交道
import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// := 符号用于短变量声明
	// os.Args是一个字符串slice,动态容量的顺组数组，可以通过s[i]来访问单个元素
	for i :=1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
