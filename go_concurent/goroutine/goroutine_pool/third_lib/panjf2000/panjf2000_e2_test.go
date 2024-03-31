package panjf2000

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"math/rand"
	"runtime"
	"sync"
	"testing"
)

const (
	DataSize = 10000
	DataPerTask = 100
)

func TestPanjf2000V2(t *testing.T)  {
	numCPUs := runtime.NumCPU()

	p := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		var sum int
		for _, n := range payload.([]int){
			sum +=n
		}
		return sum
	})
	defer p.Close()

	// 生成测试数据，还是 10000 个随机数，分成 100 组：
	nums := make([]int, DataSize)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}

	// tunny与ants不同的是，tunny的任务处理是同步的，即调用p.Process()方法之后，当前 goroutine 会挂起，直到任务处理完成之后才会被唤醒。
	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)
	partialSums := make([]int, DataSize/DataPerTask)
	for i := 0; i < DataSize/DataPerTask; i++ {
		go func(i int) {
			partialSums[i] = p.Process(nums[i*DataPerTask : (i+1)*DataPerTask]).(int)
			wg.Done()
		}(i)
	}

	wg.Wait()

	var sum int
	for _, s := range partialSums {
		sum += s
	}

	var expect int
	for _, num := range nums {
		expect += num
	}
	fmt.Printf("finish all tasks, result is %d expect:%d\n", sum, expect)
}

