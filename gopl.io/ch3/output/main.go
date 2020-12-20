package main

import "fmt"

func main() {
	medals := []string{"aa", "bb", "cc"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	// 八进制
	o := 0666
	// %后面的副词[1]告知printf重复使用第一个操作数
	// %o %x %X之前的副词#告知printf输出相应的前缀0 0x 0X
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X \n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	// %c 输出文字符号，输出带单引号 %q
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)
}
