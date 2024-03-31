package v1

import "fmt"

/**
协程池
*/
type Pool struct {
	// 工作协程的chan，无缓冲区（同步）
	work chan func()

	// 控制并发数，带缓冲区
	sem chan struct{}
}

func NewPool(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func (pool *Pool) NewTask(task func()) {
	select {
	case pool.work <- task:
		fmt.Println("pool.work sends success.")
	case pool.sem <- struct{}{}:
		go pool.worker(task)
	}
}

func (pool *Pool) worker(task func()) {
	defer func() {
		<-pool.sem //理论上是不会走这个流程
	}()

	// 重复利用开启的goroutine
	for {
		task()

		// 消费者 (如果消费者没准备好，同步的channel就不会发送成功，也就是pool.work <-task 事件不会被触发
		task = <-pool.work
	}
}

/**
优点：可最大化利用协程资源，如果任务频繁，可使用该方式

缺点：
1、开启的协程资源不会被回收，即使没有任务。
2、如果任务中包含阻塞操作，会引发其他任务无法获取资源而一直处于等待
*/
