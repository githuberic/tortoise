package e1

import (
	"fmt"
	"testing"
)

/**
all goroutines are asleep - deadlock!
*/
func TestVerify(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		defer close(naturals)
		for x := 1; x <= 100; x++ {
			naturals <- x
		}
	}()

	go func() {
		defer close(squares)
		for x := range naturals {
			squares <- x * x
		}
	}()

	// <-squares 处，期望从管道中拿到数据，而这个数据是其他goroutine协程放入管道
	// 但是其他goroutine线都已经执行完了(all goroutines are asleep)
	// 那么就永远不会有数据放入管道。所以，main goroutine线在等一个永远不会来的数据，
	// 那整个程序就永远等下去了，所以程序down掉。
	for v := range squares {
		fmt.Println(v)
	}
}
