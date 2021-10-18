package v1

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	x := 3
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x<<2)
	fmt.Printf("%08b\n", x<<3)

	y := 12
	fmt.Printf("%08b,%d\n", y, y)
	fmt.Printf("%08b,%d\n", y<<2, y<<2)
	fmt.Printf("%08b,%d\n", y<<3, y<<3)
	fmt.Printf("%08b,%d\n", y<<4, y<<4)

	z := 12
	fmt.Printf("%08b,%d\n", z, z)
	fmt.Printf("%08b,%d\n", z>>1, z>>1)
	fmt.Printf("%08b,%d\n", z>>2, z>>2)
}
