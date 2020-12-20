package main

/**
 指针对于flag，使用程序命令行参数来设置整个程序内某些变量的值
 */
import (
	"flag"
	"fmt"
	"strings"
)

// flag.Bool  函数创建一个新的boolean标识变量，三个参数 标识的名字(n),变量的默认值false,以及当用户提供非法标识、非法参数、-h/-help参数时输出的信息
// 变量sep/n是指向标识变量的指针，必须通过*set/*n来访问
var n = flag.Bool("n",false,"omit trailing newline")
var sep = flag.String("s"," ","separator")

func main()  {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n {
		fmt.Println()
	}
}
