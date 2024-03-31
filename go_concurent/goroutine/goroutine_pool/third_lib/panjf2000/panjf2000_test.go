package panjf2000

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"testing"
	"time"
)

func worker(i int, wg *sync.WaitGroup) func() {
	var cnt int

	return func() {
		for {
			time.Sleep(time.Millisecond * 200)
			cnt++
			fmt.Printf("worker processed %dth job\n", i)
			//fmt.Sprintf("worker %d processed job: %d, it's the %dtd processed by me", index, job, cnt)


			if cnt > 5 && i == 1 {
				fmt.Println("Exit go-routine id:", i)
				break
			}
		}
		wg.Done()
	}
}

func TestPanjf2000(t *testing.T) {
	wg := sync.WaitGroup{}

	pool, _ := ants.NewPool(3)
	defer pool.Release()

	for i := 1; i <= 5; i++ {
		pool.Submit(worker(i, &wg))
		wg.Add(1)
	}

	wg.Wait()
}
