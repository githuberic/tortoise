package go_bad

import (
	"fmt"
	"testing"
)

func fibonacci(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func TestGoRoutineLeakFibonacci(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)

	go fibonacci(c, quit)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	//close(quit)
}
