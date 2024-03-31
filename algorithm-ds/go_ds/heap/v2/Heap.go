package v2

import "algorithm-ds/go_ds/heap"

type HeapV2 struct {
	heap.Heap
}

func NewHeapV2(arr []int) *HeapV2 {
	h := heap.NewHeap(arr)

	heapV1 := HeapV2{}
	heapV1.Heap = *h
	return &heapV1
}

func (h *HeapV2) down(i int) {
	//存储u
	tmp := i
	n := len(h.Arr)
	left := 2*i + 1
	if left < n && h.Less(left, tmp) {
		tmp = left
	}

	right := 2*i + 2
	if right < n && h.Less(right, tmp) {
		tmp = right
	}

	if tmp != i {
		h.Less(i, tmp)
		h.down(tmp)
	}
}
