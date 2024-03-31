package sort

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	arr := []int{20, 7, 3, 10, 15, 25, 30, 17, 19}

	sortedArr := Sort(arr)
	fmt.Printf("Sorted %+v\n", sortedArr)
}
