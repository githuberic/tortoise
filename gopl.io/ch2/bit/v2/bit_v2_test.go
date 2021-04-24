package v2

import (
	"fmt"
	"math/rand"
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


	isOddNumber();
}

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
