package go_interesting_code

import (
	"fmt"
	"testing"
)

func TestRandomChannel(t *testing.T) {
	for i := range random(100) {
		fmt.Println(i)
	}
}

/**
基于管道的随机数生成器

随机数的一个特点是不好预测。如果一个随机数的输出是可以简单预测的，那么一般会称为伪随机数。
*/
func random(n int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	return ch
}
