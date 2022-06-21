package e1

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	slice := make([]int, 3, 5)
	fmt.Println("before", slice)
	changeSliceMember(slice)
	fmt.Println("alter", slice)
}

func changeSliceMember(slice []int) {
	if len(slice) > 0 {
		slice[0] = 88
	}
}
