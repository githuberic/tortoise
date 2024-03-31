package e3

import (
	"fmt"
	"testing"
)

func counter(out chan<- int) {
	for i := 1; i <= 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func TestVerify(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)

	printer(squares)
}
