package go_mem

import (
	"fmt"
	"testing"
)

func TestVerifyV4(t *testing.T) {
	s1 := new([]int)
	fmt.Println(s1)
	fmt.Printf("%p\n", s1)

	s2 := make([]int, 10)
	s2[1] = 1
	fmt.Println(s2)
	fmt.Printf("%p\n",&s2)
}
