package sort

import (
	"algorithm-ds/go_ds/heap/v1"
)

func Sort(arr []int) []int {
	heapV1 := v1.NewHeapV1(arr)
	heapV1.Heapify()

	var res []int
	for len(heapV1.Arr) != 0 {
		res = append(res, heapV1.Pop())
	}
	return res
}
