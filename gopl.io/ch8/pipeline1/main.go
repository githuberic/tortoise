package main

import "fmt"

/**
all goroutines are asleep - deadlock!
 */
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
		close(squares)
	}()

	// <-squares 处，期望从管道中拿到数据，而这个数据是其他goroutine协程放入管道
	// 但是其他goroutine线都已经执行完了(all goroutines are asleep)
	// 那么就永远不会有数据放入管道。所以，main goroutine线在等一个永远不会来的数据，
	// 那整个程序就永远等下去了，所以程序down掉。
	for {
		fmt.Println(<-squares)
	}
}
