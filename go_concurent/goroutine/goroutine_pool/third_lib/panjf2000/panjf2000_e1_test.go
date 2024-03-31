package panjf2000

import (
	"github.com/panjf2000/ants/v2"
	"log"
	"sync"
	"time"
)

const (
	runTime = 100
)

func antsDefaultCommon() {
	// 这里使用等待是为了看出结果，阻塞主线程，防止直接停止，如果在web项目中，就不需要
	var wg sync.WaitGroup

	// 退出工作，相当于使用后关闭
	defer ants.Release()
	log.Println("start ants work")
	for i := 0; i < runTime; i++ {
		wg.Add(1)
		ants.Submit(func() {
			// 提交函数，将逻辑函数提交至work中执行，这里写入自己的逻辑
			log.Println(i, ":hello")
			time.Sleep(time.Millisecond * 10)
			wg.Done()
		})
	}
	wg.Wait()
	log.Println("stop ants work")
}

