package panjf2000

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func TestTimeOut(t *testing.T)  {
	numCPUs := runtime.NumCPU()

	p := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		n := payload.(int)
		result := fib(n)
		//fmt.Println("fib() result=",result)
		time.Sleep(1000 * time.Millisecond)
		return result
	})
	defer p.Close()

	var wg sync.WaitGroup
	wg.Add(numCPUs)
	for i := 0; i < numCPUs; i++ {
		go func(i int) {
			n := rand.Intn(30)
			result, err := p.ProcessTimed(n, time.Second)
			nowStr := time.Now().Format("2006-01-02 15:04:05")
			if err != nil {
				fmt.Printf("[%s]task(%d) failed:%v\n", nowStr, i, err)
			} else {
				fmt.Printf("[%s]fib(%d) = %d\n", nowStr, n, result)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}



