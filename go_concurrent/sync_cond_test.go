package go_concurrent

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
	time.Sleep(time.Second * 1)
	cond.L.Unlock()
}

func TestVerifyV4(t *testing.T) {
	for i := 0; i < 10; i++ {
		go lockMethod(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 3)
	fmt.Println("signal1")

	cond.Signal()
	time.Sleep(time.Second * 1)
	fmt.Println("signal2")

	cond.Signal()
	time.Sleep(time.Second * 1)

	fmt.Println(">>>start broadcast")
	cond.Broadcast()//下发广播给所有等待的goroutine
	time.Sleep(time.Second * 60)
	fmt.Println(">>>end broadcast")
}
