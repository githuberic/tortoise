package go_good

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func fibonacci(ch chan int, ctx context.Context) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-ctx.Done():
			fmt.Println("quit")
			return
		}
	}
}

func TestGoRoutineLeakFibonacci(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)

	go fibonacci(ch, ctx)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	cancel()

	time.Sleep(time.Second * 5)
}
