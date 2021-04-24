package v2

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestVerify(t *testing.T) {
	var x int32 = 0xAC
	fmt.Printf("%d,%08b\n", x, x)
	// AND 操作符是一个很好的将整数的指定位清零的方式,使用 & 运算符将数字后 4 位清零。
	x = x & 0xF0
	fmt.Printf("%d,%08b\n", x, x)

	x &= 0xF0
	fmt.Printf("%d,%08b\n", x, x)

	xorTest()
	//orTest()

	//fmt.Println(procstr("HELLO World!", LOWER|REV|CAP))
}

/**
& 来判断一个数字是奇数还是偶数。我们可以将数字和值 1 使用 & 做 AND 运算。
如果结果是 1，那说明原来的数字是一个奇数。
*/
func isOddNumber() {
	for x := 0; x <= 100; x++ {
		num := rand.Int()
		if num&1 == 1 {
			fmt.Printf("%d is odd\n", num)
		} else {
			fmt.Printf("%d is even\n", num)
		}
	}
}

/**
可以使用这个特性来将一个整数中的指定位置为 1。
在下面的例子里，我们使用 OR 运算将第 3、7、8 位置为 1。
*/
func orTest() {
	var x int64 = 0
	x |= 196
	fmt.Printf("%d,%08b\n", x, x)

	// 数字使用掩码技术，OR 操作
	x |= 3
	fmt.Printf("%d,%08b\n", x, x)
}

const (
	UPPER = 1 // upper case
	LOWER = 2 // lower case
	CAP   = 4 // capitalizes
	REV   = 8 // reverses
)

func procstr(str string, conf byte) string {
	// reverse string
	rev := func(s string) string {
		runes := []rune(s)
		n := len(runes)
		for i := 0; i < n/2; i++ {
			runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
		}
		return string(runes)
	}

	// query config bits
	if (conf & UPPER) != 0 {
		str = strings.ToUpper(str)
	}
	if (conf & LOWER) != 0 {
		str = strings.ToLower(str)
	}
	if (conf & CAP) != 0 {
		str = strings.Title(str)
	}
	if (conf & REV) != 0 {
		str = rev(str)
	}
	return str
}

func xorTest() {
	var x uint64 = 0xCEFF
	fmt.Printf("%d,%08b\n", x, x)
	x ^= 0xFF00
	fmt.Printf("%d,%08b\n", x, x)
}
