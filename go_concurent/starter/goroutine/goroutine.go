package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	TIMES = 10 * 1000 * 1000
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() + 1)
	fmt.Println("CPUs:", runtime.NumCPU(), "Goroutines:", runtime.NumGoroutine())

	t1 := time.Now()
	for i := 0; i < TIMES; i++ {
		go func() {}()
	}

	t2 := time.Now()
	fmt.Printf("elapsed time: %.3fs\n", t2.Sub(t1).Seconds())
}
