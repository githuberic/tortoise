## Channel 同步

```go
package sync

import (
	"fmt"
	"testing"
	"time"
)
// 这个worker函数将以协程的方式运行
// 通道`done`被用来通知另外一个协程这个worker函数已经执行完成
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 1)
	fmt.Print("done")
	// 向通道发送一个数据，表示worker函数已经执行完成
	done <- true
}

func TestVerify(t *testing.T) {
	// 使用协程来调用worker函数，同时将通道`done`传递给协程
	// 以使得协程可以通知别的协程自己已经执行完成
	var done = make(chan bool, 1)
	go worker(done)

	// 一直阻塞，直到从worker所在协程获得一个worker执行完成的数据
	<-done
}
```

这是一个我们将要在 Go 协程中运行的函数。done 通道将被用于通知其他 Go 协程这个函数已经工作完毕。

发送一个值来通知我们已经完工啦。

运行一个 worker Go协程，并给予用于通知的通道。

程序将在接收到通道中 worker 发出的通知前一直阻塞。

如果你把 <- done 这行代码从程序中移除，程序甚至会在 worker还没开始运行时就结束了。