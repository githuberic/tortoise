package _select

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func doGoodthing(done chan bool) {
	time.Sleep(time.Second * 1)

	select {
	case done <- true:
	default:
		return
	}
}

func timeout(f func(done chan bool)) error {
	done := make(chan bool)

	go f(done)

	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func test(t *testing.T, f func(chan bool)) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		timeout(f)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func TestBadTimeout(t *testing.T) {
	test(t, doGoodthing)
}
