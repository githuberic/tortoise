package sort

import (
	"fmt"
	"testing"
)

var hs *HeapSortV1

func init() {
	arr := []int{8, 10, 2, 4, 3, 5, 1, 6, 7}
	hs = NewHeapSortV1(arr)
	hs.buildHeap(arr)
}

func TestHeapBuild(t *testing.T)  {
	fmt.Printf("Build heap:%v,\n", hs.Arr)
}

func TestHeapSort(t *testing.T) {
	arr := []int{8, 10, 2, 4, 3, 5, 1, 6, 7}
	hs.Sort(arr)
	fmt.Printf("Heap sort:%v,\n", hs.Arr)
}
