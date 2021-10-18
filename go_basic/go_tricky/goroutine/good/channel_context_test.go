package go_good

import (
	"context"
	"fmt"
	"testing"
)

/**
我们可以通过context包来避免这个问题：
当main函数在break跳出循环时，通过调用cancel()来通知后台Goroutine退出，这样就避免了Goroutine的泄漏。
*/
func TestChannelReturn(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			cancel()
			break
		}
	}
}
