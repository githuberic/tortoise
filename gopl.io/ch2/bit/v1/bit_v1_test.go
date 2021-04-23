package v1

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	var x int64 = 3
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x<<2)
	fmt.Printf("%08b\n", x<<3)

	var y int = 12
	fmt.Printf("%08b\n", y<<2)
	fmt.Printf("%d\n", y<<2)

	fmt.Printf("%08b\n", y<<3)
	fmt.Printf("%d\n", y<<3)

	fmt.Printf("%08b\n", y<<4)
	fmt.Printf("%d\n", y<<4)
}
