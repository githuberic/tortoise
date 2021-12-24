package go_pointer

import (
	"fmt"
	"testing"
)

/**
当一个指针被定义后没有分配到任何变量时，它的值为 nil
*/
func TestNullPointer(t *testing.T) {
	var p *string
	fmt.Println(p)
	fmt.Printf("P value = %v\n", p)
	if p == nil {
		fmt.Println("Value is null")
	} else {
		fmt.Println("Value not null")
	}
}
