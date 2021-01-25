package go_mem

import (
	"fmt"
	"testing"
)

func TestVerifyV2(t *testing.T) {
	// cap不够时，cap < 1024, cap容量成倍增加；cap >= 1024时，按照1.25倍扩容
	arr := []int{1}
	fmt.Printf("len: %d cap:%d data:%+v\n", len(arr), cap(arr), arr)

	arr = append(arr, 2)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(arr), cap(arr), arr)

	arr = append(arr, 3)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(arr), cap(arr), arr)

	arr = append(arr, 4, 5)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(arr), cap(arr), arr)

	a := []int{1, 2, 3}
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)

	appendSlice(a)
	// 4去哪儿了？
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a), cap(a), a)

	a2 := appendSliceV2(a)
	// 4去哪儿了？
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a2), cap(a2), a2)

	appendSliceV3(&a)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a2), cap(a2), a2)
}

func appendSlice(arr []int) {
	arr = append(arr, 4)
}

func appendSliceV2(arr []int) []int {
	arr = append(arr, 4)
	return arr
}

// 指针传参
func appendSliceV3(arr *[]int) {
	*arr = append(*arr, 4)
}
