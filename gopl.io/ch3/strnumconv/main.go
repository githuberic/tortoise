package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 123
	var y string = fmt.Sprintf("%d", x)
	fmt.Println(y)
	fmt.Println(strconv.Itoa(x))
	// FormatInt/FormatUint 可以按不同位进制格式化数字
	fmt.Println(strconv.FormatInt(int64(x), 2))
	var str string = fmt.Sprintf("x=%b", x)
	fmt.Println(str)

	// strconv 包内的atoi/parseInt用于解析整数的字符串
	x, err := strconv.Atoi("123")
	if err == nil {
		fmt.Println(x)
	}
	// 十进制,最长为64位
	x1, err := strconv.ParseInt("123", 10, 64)
	if err == nil {
		fmt.Println(x1)
	}
}
