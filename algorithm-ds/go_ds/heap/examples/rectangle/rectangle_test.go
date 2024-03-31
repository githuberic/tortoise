package rectangle

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestRectangle(t *testing.T) {
	hp := &RectHeap{}
	for i := 2; i < 6; i++ {
		*hp = append(*hp, Rectangle{i, i})
	}

	fmt.Println("原始slice: ", hp)

	// 堆操作
	heap.Init(hp)
	heap.Push(hp, Rectangle{100, 10})
	fmt.Println("Top ele：", (*hp)[0])
	fmt.Println("Remove,return last ele：", heap.Pop(hp)) // 最后 一个元素
	fmt.Println("Lase slice: ", hp)
}
