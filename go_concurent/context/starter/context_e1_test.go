package starter

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func doTask(n int) {
	time.Sleep(time.Duration(n))
	fmt.Printf("Task %d Done\n", n)
	wg.Done()
}

func TestVerify(t *testing.T) {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doTask(i+1)
	}
	wg.Wait()
	fmt.Println("All Task Done")
}
