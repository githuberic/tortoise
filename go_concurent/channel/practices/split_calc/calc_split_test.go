package split_calc

import (
	"fmt"
	"testing"
	"time"
)

func TestSplitCalc(t *testing.T) {
	var result, i uint64
	start := time.Now()
	for i = 1; i <= 10000000000; i++ {
		result += i
	}
	elapsed := time.Since(start)
	fmt.Println("执行消耗时间为:", elapsed)
	fmt.Println(",result:", result)

	// 4个协程共同执行累加操作
	start = time.Now()
	ch1 := calc(1, 2500000000)
	ch2 := calc(2500000001, 5000000000)
	ch3 := calc(5000000001, 7500000000)
	ch4 := calc(7500000001, 10000000000)

	// 汇总4个协程的累加结果
	result = <-ch1 + <-ch2 + <-ch3 + <-ch4
	// 统计计算耗时
	elapsed = time.Since(start)
	fmt.Println("执行消耗的时间为:", elapsed)
	fmt.Println(", result:", result)
}

func calc(from, to uint64) <-chan uint64 {
	ch := make(chan uint64)
	go func() {
		result := from
		for i := from + 1; i <= to; i++ {
			result += i
		}
		ch <- result
	}()
	return ch
}
