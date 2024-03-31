package rxgo

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

/**
ForEach()接受 3 个回调函数：
	NextFunc：类型为func (v interface {})，处理数据；
	ErrFunc：类型为func (err error)，处理错误；
	CompletedFunc：类型为func ()，Observable 完成时调用。

ForEach()实际上是异步执行的，它返回一个接收通知的 channel。
当 Observable 数据发送完毕时，该 channel 会关闭。
所以如果要等待ForEach()执行完成，我们需要使用<-。
上面的示例中如果去掉<-，可能就没有输出了，因为主 goroutine 结束了，整个程序就退出了。
*/
func TestRxgoJustForEach(t *testing.T) {
	observable := rxgo.Just(1, 2, 3, 4, 5)()
	<-observable.ForEach(func(v interface{}) {
		fmt.Println("Received:", v)
	}, func(err error) {
		fmt.Println("error:", err)
	}, func() {
		fmt.Println("completed")
	})
}
