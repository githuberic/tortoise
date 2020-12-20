package main

import (
	"fmt"
	"unicode/utf8"
)

/**
字符字节数组
*/
func main() {
	var str string = "hello, world"
	fmt.Println(len(str))
	fmt.Print(str[0], str[7])
	fmt.Println(str[0:5])
	fmt.Println(str[:5])
	fmt.Println(str[7:])

	strTemp := "hello, 世界"
	fmt.Println(len(strTemp))
	fmt.Println(utf8.DecodeRuneInString(strTemp))

	for i := 0; i < len(strTemp); {
		// DecodeRuneInString 每次调用都返回r(文字符号本身)/值(r表示按utf-8编码所占用的字节数)
		r, size := utf8.DecodeRuneInString(strTemp[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	// range循环内隐式读取
	for i, r := range "hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
