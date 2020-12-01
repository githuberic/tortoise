package main

/**
 指针对于flag，使用程序命令行参数来设置整个程序内某些变量的值
 */
import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n",false,"omit trailing newline")
var sep = flag.String("s"," ","separator")

func main()  {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n {
		fmt.Println()
	}
}
