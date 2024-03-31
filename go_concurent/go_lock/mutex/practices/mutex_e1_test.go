package practices

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestVerify(t *testing.T) {
	// 这个例子的状态就是一个map
	var state = make(map[int]int)

	// 这个`mutex`将同步对状态的访问
	//var mutex = &sync.Mutex{}
	var mutex sync.Mutex

	// ops将对状态的操作进行计数
	var ops int64 = 0

	// 这里我们启动100个协程来不断地读取这个状态
	for r := 0; r < 100; r++ {
		go func() {
			var total int = 0
			for {
				// 对于每次读取，我们选取一个key来访问，
				// mutex的`Lock`函数用来保证对状态的
				// 唯一性访问，访问结束后，使用`Unlock`
				// 来解锁，然后增加ops计数器
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	// 同样我们使用10个协程来模拟写状态
	for w := 0; w < 10; w++ {
		go func() {
			for {
				var key = rand.Intn(5)
				var val = rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	// 主协程Sleep，让那10个协程能够运行一段时间
	time.Sleep(time.Second * 1)

	// 输出总操作次数
	var opsFinal = atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	// 最后锁定并输出状态
	mutex.Lock()
	fmt.Println("State:", state)
	mutex.Unlock()
}
