package v1

import (
	"fmt"
	"testing"
	"time"
)

func producer(arr []interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer func() {
			close(out)
			fmt.Println("Producer exit")
		}()

		for _, v := range arr {
			fmt.Printf("send %d\n", v)
			out <- v
			time.Sleep(time.Millisecond)
		}
	}()
	return out
}

func consumer(in <-chan interface{}) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer func() {
			fmt.Println("consumer exit")
			done <- struct{}{}
			close(done)
		}()

		for v := range in{
			fmt.Printf("Processing %d\n",v)
		}
	}()
	return done
}

func TestDetectClose(t *testing.T) {
	arr := []interface{}{3, 4, 5, 6, 8, 9}
	out := producer(arr)
	done := consumer(out)
	<-done
	fmt.Println("Main exit")
}
