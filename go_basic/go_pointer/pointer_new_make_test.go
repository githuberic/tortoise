package go_pointer

import (
	"fmt"
	"testing"
)

func TestNewMakeV1(t *testing.T) {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)
}

func TestNewMakeV2(t *testing.T) {
	var a *int
	a = new(int)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)
}
