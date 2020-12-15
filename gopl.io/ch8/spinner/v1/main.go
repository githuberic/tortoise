package main

import (
	"fmt"
	"time"
)

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration, str string) {
	for {
		fmt.Printf("%s", str)
		time.Sleep(delay)
	}
}

func main() {
	fmt.Println(">>>Start")
	go spinner(100*time.Millisecond, ".")
	const n = 45
	finN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, finN)
}
