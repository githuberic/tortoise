package v2

import "fmt"

type Pool2 struct {
	num chan struct{}
}

func NewPool2(size int) *Pool2 {
	return &Pool2{
		num: make(chan struct{}, size),
	}
}

func (pool *Pool2) NewTask(task func()) {
	select {
	case pool.num <- struct{}{}:
		go pool.Worker(task)
	}
}

func (pool *Pool2) Worker(task func()) {
	defer func() {
		fmt.Println("go routing ends.")
		<-pool.num
	}()

	// 此处不能使用go 开启协程，如果开启go，那么defer的内容会被直接执行，也就是达不到限制同时使用的goroutine 数量
	task()
}

/**
与第一种方法一样，使用chan来控制并发数，但是不需要work chan，因为每次都是新建协程，在未达到最大并发之前，直接执行即可。
*/
