package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	// map 存储一个键/值对集合
	counts := make(map[string]int)

	// bufio 高效的输入和输出
	input := bufio.NewScanner(os.Stdin)

	// 扫描器从程序的标准输入进行读取，每一次调用input.Scan()读取下一行，并且将结尾的换行符去掉，通过input.text读取内容
	for  input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n",n ,line)
		}
	}
}
