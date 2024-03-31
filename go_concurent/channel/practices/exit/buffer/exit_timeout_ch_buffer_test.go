package buffer

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func doBadthing(done chan bool) {
	time.Sleep(time.Second * 1)
	done <- true
}
func timeoutWithBuffer(f func(done chan bool)) error {
	done := make(chan bool, 1)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func TestBufferTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		timeoutWithBuffer(doBadthing)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}
