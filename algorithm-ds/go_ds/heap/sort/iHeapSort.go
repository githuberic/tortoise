package sort

// 堆排序接口
type iHeapSort interface {
	// 堆化
	Heapify(i int)

	// 排序
	Sort(arr []int)
}
