package sort

import "algorithm-ds/go_ds/heap"

type HeapSortV1 struct {
	heap.Heap
}

// 本实例构建最小堆 /**
func NewHeapSortV1(arr []int) *HeapSortV1 {
	h := heap.NewHeap(arr)

	heapV1 := HeapSortV1{}
	heapV1.Heap = *h
	return &heapV1
}

// 调整堆,第i个结点
func (h *HeapSortV1) Heapify(i int) {
	n := len(h.Arr)
	if i >= n {
		return
	}

	// left 左子结点,right右左子结点
	left := 2*i + 1
	right := 2*i + 2
	max := i

	if left < n && h.Greater(left, max) {
		max = left
	}
	if right < n && h.Greater(right, max) {
		max = right
	}

	if max != i {
		h.Swap(max, i)
		h.Heapify(max)
	}
}

// 创建堆(构建最大堆)
func (h *HeapSortV1) buildHeap(arr []int) {
	lastNode := len(arr) - 1     // 最后一个结点
	parent := (lastNode - 1) / 2 // 最后一个结点的父节点

	// 往最后一个父节点开始，不断往前创建结点
	for i := parent; i >= 0; i-- {
		h.Heapify(i)
	}
}

func (h *HeapSortV1) Sort(arr []int) {
	h.buildHeap(arr)
	n := len(arr)

	// 重复：第一个结点和最后一个结点交换位置，然后重新调整堆排序，i--去掉最后一个结点
	for i := n - 1; i >= 0; i-- {
		h.Swap(i, 0)
		h.Heapify(i)
	}
}
