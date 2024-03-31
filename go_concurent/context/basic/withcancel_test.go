package basic

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func doTaskE3(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func TestVerifyE3(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go doTaskE3(ctx,"Work1")
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

func TestVerifyE4(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go doTaskE3(ctx,"Work1")
	go doTaskE3(ctx,"Work2")
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
