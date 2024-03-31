package v1

import (
	"fmt"
	"testing"
)

var h *HeapV1

func init() {
	arr := []int{20, 7, 3, 10, 15, 25, 30, 17, 19}
	h = NewHeapV1(arr)
	h.Heapify()
}

func TestHeapInit(t *testing.T) {
	fmt.Println("Heapify", h.String()) // [3 7 20 10 15 25 30 17 19]
}

func TestHeapPush(t *testing.T) {
	h.Push(6)
	fmt.Println("Pushed ", h.String())
}

func TestHeapRemove(t *testing.T) {
	x, ok := h.Remove(5)
	fmt.Println("Removed=", x, "Result=", ok, h.String()) // 25 true [3 6 15 10 7 20 30 17 19]
}

func TestHeap(t *testing.T) {
	z := h.Pop()
	fmt.Println("Pop=", z, h.String())
}
