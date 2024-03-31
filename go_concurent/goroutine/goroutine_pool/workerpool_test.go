package goroutine_pool

import (
	"fmt"
	"testing"
	"time"
)

func worker(index int, jobCh <-chan int, retCh chan<- string) {
	cnt := 0
	for job := range jobCh {
		cnt++
		ret := fmt.Sprintf("worker %d processed job: %d, it's the %dtd processed by me", index, job, cnt)
		retCh <- ret
	}
}

func workerPool(count int, jobCh <-chan int, retCh chan<- string) {
	for i := 0; i < count; i++ {
		go worker(i, jobCh, retCh)
	}
}

func geneJob(count int) <-chan int {
	jobCh := make(chan int, 200)

	go func() {
		for i := 0; i < count; i++ {
			jobCh <- i
		}
		close(jobCh)
	}()

	return jobCh
}

func TestWorkerPool(t *testing.T) {
	jobCh := geneJob(10000)
	retCh := make(chan string, 10000)
	workerPool(5, jobCh, retCh)

	time.Sleep(1 * time.Second)
	close(retCh)
	for ret := range retCh {
		fmt.Println(ret)
	}
}
