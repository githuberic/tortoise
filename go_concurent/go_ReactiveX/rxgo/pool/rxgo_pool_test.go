package pool

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"math/rand"
	"testing"
	"time"
)

/**
即只有一个 goroutine 负责执行转换函数。我们也可以使用rxgo.WithPool(n)选项设置运行n个 goroutine，
或者rxgo.WitCPUPool()选项设置运行与逻辑 CPU 数量相等的 goroutine。
 */
func TestPool(t *testing.T)  {
	observable := rxgo.Range(1, 100)

	observable = observable.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		time.Sleep(time.Duration(rand.Int31()))
		return i.(int)*2 + 1, nil
	}, rxgo.WithCPUPool())

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
