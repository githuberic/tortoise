package until_all_done

import (
	"fmt"
	"runtime"
	"testing"
)

func runTask(id int) string {
	//time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	finalRet := ""
	// 等待所有的结果执行完毕(此时 channel 里面 内容别读取完毕了，其goroutine释放了)
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}

	return finalRet
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	//time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
