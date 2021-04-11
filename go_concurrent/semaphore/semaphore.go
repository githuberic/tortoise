package semaphore

import "sync"

// Semaphore 数据结构，并且还实现了Locker接口
type semaphore struct {
	sync.Locker
	ch chan struct{}
}

// 创建一个新的信号量
func (*semaphore) NewSemaphore(capacity int) sync.Locker {
	if capacity <= 0 {
		// 容量为1就变成了一个互斥锁
		capacity = 1
	}
	return &semaphore{ch: make(chan struct{}, capacity)}
}

// 请求一个资源
func (p *semaphore) Lock() {
	p.ch <- struct{}{}
}

// 释放资源
func (p *semaphore) Unlock() {
	<-p.ch
}
