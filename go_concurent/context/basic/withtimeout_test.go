package basic

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func doTaskE5(ctx context.Context, name string) {
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

func TestVerifyE6(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	go doTaskE5(ctx, "worker1")
	go doTaskE5(ctx, "worker2")
	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel()
	time.Sleep(3 * time.Second)
}


