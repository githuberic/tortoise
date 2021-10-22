package heap


//build a heap
func buidHeap(a []int, n int) {

	//heapify from the last parent node
	for i := n / 2; i >= 1; i-- {
		heapifyUpToDown(a, i, n)
	}
}

//sort by ascend, a index begin from 1, has n elements
func sort(a []int, n int) {
	buidHeap(a, n)

	k := n
	for k >= 1 {
		swap(a, 1, k)
		heapifyUpToDown(a, 1, k-1)
		k--
	}
}