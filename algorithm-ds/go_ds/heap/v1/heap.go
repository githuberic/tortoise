package v1

import "algorithm-ds/go_ds/heap"

type HeapV1 struct {
	heap.Heap
}

// 本实例构建最小堆 /**
func NewHeapV1(arr []int) *HeapV1 {
	h := heap.NewHeap(arr)

	heapV1 := HeapV1{}
	heapV1.Heap = *h
	return &heapV1
}

func (h *HeapV1) Push(x int) {
	h.Arr = append(h.Arr, x)
	h.up(len(h.Arr) - 1)
}

func (h *HeapV1) Remove(i int) (int, bool) {
	if i < 0 || i > len(h.Arr)-1 {
		return 0, false
	}

	n := len(h.Arr) - 1
	h.Swap(i, n) // 用最后的元素值替换被删除元素
	// 删除最后的元素
	x := (h.Arr)[n]
	h.Arr = (h.Arr)[0:n]

	// 如果当前元素大于父结点，向下筛选
	if (h.Arr)[i] > (h.Arr)[(i-1)/2] {
		h.down(i)
	} else {
		// 当前元素小于父结点，向上筛选
		h.up(i)
	}

	return x, true
}

// Pop 弹出堆顶的元素，并返回其值 /**
func (h *HeapV1) Pop() int {
	n := len(h.Arr) - 1
	h.Swap(0, n)
	x := (h.Arr)[n]
	h.Arr = (h.Arr)[0:n]
	h.down(0)
	return x
}

func (h *HeapV1) Heapify() {
	n := len(h.Arr)

	// i > n/2-1 的结点为叶子结点本身已经是堆了
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *HeapV1) up(i int) {
	for {
		// 父亲结点
		parent := (i - 1) / 2
		if i == parent || h.Less(parent, i) {
			break
		}

		// 如果父结点比孩子结点小，则交换
		h.Swap(parent, i)
		// 继续向下比较
		i = parent
	}
}

func (h *HeapV1) down(i int) {
	for {
		// Left
		left := 2*i + 1
		if left >= len(h.Arr) {
			break // i已经是叶子结点了
		}

		// 从left,right中选择min,
		j := left
		if r := left + 1; r < len(h.Arr) && h.Less(r, left) {
			j = r // 右孩子
		}

		if h.Less(i, j) {
			break // 如果父结点比孩子结点小，则不交换
		}

		h.Swap(i, j) // 交换父结点和子结点
		i = j        //继续向下比较
	}
}
