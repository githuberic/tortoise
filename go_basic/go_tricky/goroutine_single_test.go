package go_tricky

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

// Goroutine是协作式抢占调度，Goroutine本身不会主动放弃CPU
func TestGoRoutine(t *testing.T) {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	// 占用CPU
	select {}
}

// 解决的方法是在for循环加入runtime.Gosched()调度函数：
func TestGoRoutineV2(t *testing.T) {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	for {
		runtime.Gosched()
	}
}

// 或者是通过阻塞的方式避免CPU占用：
func TestGoRoutineV3(t *testing.T) {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		os.Exit(0)
	}()

	for {}
}
