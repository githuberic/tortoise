package cond

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func lockMethod(x int) {
	cond.L.Lock()
	cond.Wait()
	fmt.Println(x)
	time.Sleep(time.Millisecond * 100)
	cond.L.Unlock()
}

func TestVerifyV4(t *testing.T) {
	for i := 0; i < 10; i++ {
		go lockMethod(i)
	}

	fmt.Println("Start all")
	time.Sleep(time.Millisecond * 200)
	fmt.Println("signal-1")
	cond.Signal()

	time.Sleep(time.Millisecond * 200)
	fmt.Println("signal-2")
	cond.Signal()

	time.Sleep(time.Millisecond * 200)

	fmt.Println(">>>Start broadcast")
	cond.Broadcast() //下发广播给所有等待的goroutine
	time.Sleep(time.Second * 5)
	fmt.Println(">>>End broadcast")
}
