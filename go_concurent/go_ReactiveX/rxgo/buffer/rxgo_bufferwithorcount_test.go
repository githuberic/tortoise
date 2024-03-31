package buffer

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
	"time"
)

func TestBufferWithTimerOrCount(t *testing.T) {
	ch := make(chan rxgo.Item, 1)

	go func() {
		i := 0
		for range time.Tick(time.Second * 1) {
			ch <- rxgo.Of(i)
			i++
		}
	}()

	// 3s 可以收集 3 个数据，但是设置了收集 2 个就发送
	observable := rxgo.FromChannel(ch).BufferWithTimeOrCount(rxgo.WithDuration(3*time.Second), 2)

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
