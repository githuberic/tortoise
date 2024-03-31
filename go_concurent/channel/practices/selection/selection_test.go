package selection

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	// 本例中，我们从两个通道中选择
	var ch1 = make(chan string, 1)
	var ch2 = make(chan string, 2)

	// 为了模拟并行协程的阻塞操作，我们让每个通道在一段时间后再写入一个值
	go func() {
		//time.Sleep(time.Second * 1)
		ch1 <- "one"
	}()

	go func() {
		//time.Sleep(time.Second * 1)
		ch2 <- "two"
	}()

	// 我们使用select来等待这两个通道的值，然后输出
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		}
	}
}
