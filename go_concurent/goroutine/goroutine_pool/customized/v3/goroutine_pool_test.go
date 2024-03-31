package v3

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutinePool(t *testing.T) {
	task := NewTask(func() error {
		fmt.Println("Create task,", time.Now().Format("2006-01-02 15:01:05"))
		return nil
	})

	pool := NewPool(10)

	//开一个协程 不断的向 Pool 输送打印一条时间的task任务
	go func() {
		for {
			pool.EntryChannel <- task
		}
	}()

	// 启动协程池p
	pool.Run()
}

