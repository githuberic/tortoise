package v2

import "time"

type Semaphore struct {
	permits  int           // 许可数量
	channels chan struct{} // 通道
}

/* 创建信号量 */
func NewSemaphore(permits int) *Semaphore {
	return &Semaphore{channels: make(chan struct{}, permits), permits: permits}
}

/* 获取许可 */
func (this *Semaphore) Acquire() {
	this.channels <- struct{}{}
}

/* 释放许可 */
func (this *Semaphore) Release() {
	<-this.channels
}

/* 尝试获取许可 */
func (this *Semaphore) TryAcquire() bool {
	select {
	case this.channels <- struct{}{}:
		return true
	default:
		return false
	}
}

/* 尝试指定时间内获取许可 */
func (this *Semaphore) TryAcquireOnTime(timeout time.Duration) bool {
	select {
	case this.channels <- struct{}{}:
		return true
	case <-time.After(timeout):
		return false
	}
}

/* 当前可用的许可数 */
func (this *Semaphore) AvailablePermits() int {
	return this.permits - len(this.channels)
}
