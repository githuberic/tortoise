package basic

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type Options struct {
	Interval time.Duration
}

func DoTaskE4(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			var op = ctx.Value("options").(*Options)
			time.Sleep(op.Interval * time.Second)
		}
	}
}

func TestVerifyE5(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	vCtx := context.WithValue(ctx, "options", &Options{1})

	go DoTaskE4(vCtx,"Work1")
	go DoTaskE4(vCtx,"Work2")
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
