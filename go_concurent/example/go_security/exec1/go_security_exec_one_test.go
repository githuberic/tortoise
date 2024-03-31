package exec1

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/**
【必须】禁止在闭包中直接调用循环变量
*/
func TestVerifyV1(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var group sync.WaitGroup
	for i := 1; i < 5; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			fmt.Printf("%-2d", i)
		}()
	}
	group.Wait()
}

func TestVerifyV2(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var group sync.WaitGroup

	for i := 1; i < 5; i++ {
		group.Add(1)
		go func(j int) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in start()")
				}
				group.Done()
			}()
			fmt.Printf("%-2d", j) // 闭包内部使用局部变量
		}(i) // 把循环变量显式地传给协程
	}
	group.Wait()
}
