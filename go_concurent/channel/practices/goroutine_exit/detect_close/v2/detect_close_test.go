package v2

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
		}
	}()
	return out
}

func consumer(in <-chan interface{}) <-chan struct{} {
	done := make(chan struct{})

	processedCount := 0
	t := time.NewTicker(time.Millisecond * 500)

	go func() {
		defer func() {
			fmt.Println("consumer exit")
			done <- struct{}{}
			close(done)
		}()

		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				fmt.Printf("Processing %d\n", v)
				processedCount++
			case <-t.C:
				fmt.Printf("Processing, processedCnt = %d\n", processedCount)
			}
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
