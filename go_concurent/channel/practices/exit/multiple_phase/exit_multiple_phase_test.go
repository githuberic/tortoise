package multiple_phase

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func do2Phase(phase1, done chan bool) {
	time.Sleep(time.Second * 1) // Phase1
	select {
	case phase1 <- true:
	default:
		return
	}
	time.Sleep(time.Second * 1) // Phase2
	done <- true
}

func timeoutFirstPhase() error {
	phase1 := make(chan bool)
	done := make(chan bool)
	go do2Phase(phase1, done)

	select {
	case <-phase1:
		<-done
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func Test2phasesTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		timeoutFirstPhase()
	}
	time.Sleep(time.Second * 3)
	t.Log(runtime.NumGoroutine())
}