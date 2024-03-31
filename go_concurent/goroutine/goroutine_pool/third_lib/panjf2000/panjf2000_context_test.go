package panjf2000

import (
	"context"
	"fmt"
	"github.com/Jeffail/tunny"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	numCPUs := runtime.NumCPU()
	p := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		n := payload.(int)
		result := fib(n)
		time.Sleep(2000 * time.Millisecond)
		return result
	})
	defer p.Close()

	var wg sync.WaitGroup
	wg.Add(numCPUs)

	for i := 0; i < numCPUs; i++ {
		go func(i int) {
			n := rand.Intn(30)
			ctx, cancel := context.WithCancel(context.Background())
			if i%2 == 0 {
				go func() {
					time.Sleep(500 * time.Millisecond)
					cancel()
				}()
			}

			result, err := p.ProcessCtx(ctx, n)
			if err != nil {
				fmt.Printf("task(%d) failed:%v\n", i, err)
			} else {
				fmt.Printf("fib(%d) = %d\n", n, result)
			}
			wg.Done()
		}(i)
	}
}
