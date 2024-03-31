package heap

import "fmt"

// 堆结构体
type Heap struct {
	Arr []int
}

//init heap
func NewHeap(arr []int) *Heap {
	heap := &Heap{}
	heap.Arr = arr
	return heap
}

// swap two elements
func (h *Heap) Swap(i int, j int) {
	h.Arr[i], h.Arr[j] = h.Arr[j], h.Arr[i]
}

func (h *Heap) Less(i, j int) bool {
	return h.Arr[i] < h.Arr[j]
}

func (h *Heap) Greater(i, j int) bool {
	return h.Arr[i] > h.Arr[j]
}

func (h *Heap) String() string {
	return fmt.Sprintf("Heap array:%+v", h.Arr)
}
