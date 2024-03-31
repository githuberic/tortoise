package workpool

import (
	"fmt"
	"testing"
)

// 我们将在worker函数里面运行几个并行实例，这个函数从jobs通道
// 里面接受任务，然后把运行结果发送到results通道。每个job我们
// 都休眠一会儿，来模拟一个耗时任务。
func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, ",processing job", j)
		result <- j * 2
	}
}

func TestVerify(t *testing.T) {
	// 为了使用我们的工作池，我们需要发送工作和接受工作的结果，
	// 这里我们定义两个通道，一个jobs，一个results
	var job = make(chan int, 100)
	var result = make(chan int, 100)

	// 这里启动3个worker协程，一开始的时候worker阻塞执行，因为
	// jobs通道里面还没有工作任务
	for w := 1; w <= 3; w++ {
		go worker(w, job, result)
	}

	// 这里我们发送9个任务，然后关闭通道，告知任务发送完成
	for j := 1; j <= 9; j++ {
		job <- j
	}
	close(job)

	for r := 1; r <= 9; r++ {
		<-result
	}
	close(result)
}
