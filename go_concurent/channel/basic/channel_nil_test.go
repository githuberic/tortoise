package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelNil(t *testing.T) {
	// default value of ch is nil
	var ch chan string

	go func() {
		ch = make(chan string, 1)
		ch <- "hello world"
		fmt.Println("g1 exit")
	}()

	go func(ch chan string) {
		if ch == nil {
			fmt.Println("Channel is nil")
		}
		<-ch
		fmt.Println("g2 exit")
	}(ch)

	go func() {
		fmt.Println(<-ch)
	}()
	time.Sleep(time.Second * 1)
}
