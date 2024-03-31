package e4

import (
	"fmt"
	"testing"
)

func counter(numbs ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, v := range numbs {
			out <- v
		}
	}()
	return out
}

func squarer(inCh <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range inCh {
			out <- v * v
		}
	}()
	return out
}

func printer(inCh <-chan int) {
	for v := range inCh {
		fmt.Println(v)
	}
}

func TestVerify(t *testing.T) {
	in := counter(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	outCh := squarer(in)

	printer(outCh)
}
