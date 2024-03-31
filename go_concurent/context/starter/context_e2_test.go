package starter

import (
	"fmt"
	"testing"
	"time"
)

var stop chan bool

func DoTask(name string) {
	for {
		select {
		case <-stop:
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func TestVerifyE2(t *testing.T) {
	stop = make(chan bool)
	go DoTask("Work1")
	time.Sleep(3 * time.Second)
	stop <- true
	time.Sleep(3 * time.Second)
}
