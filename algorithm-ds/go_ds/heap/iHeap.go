package heap

// 堆相关操作接口
type iHeap interface {
	Push(x int)
	Remove(i int) (int, bool)
	Pop() int
	// 堆化
	Heapify()
}
