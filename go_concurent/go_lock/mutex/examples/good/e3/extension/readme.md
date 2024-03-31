## Mutex扩展

实际上，不论是不希望锁的 goroutine 继续等待，还是想监控锁，我们都可以基于标准库中 Mutex 的实现，通过 Hacker 的方式，为 Mutex 增加一些额外的功能。这节课，我就来教你实现几个扩展功能，包括实现 TryLock，获取等待者的数量等指标，以及实现一个线程安全的队列。

### TryLock

我们可以为 Mutex 添加一个 TryLock 的方法，也就是尝试获取排外锁。

在实际开发中，如果要更新配置数据，我们通常需要加锁，这样可以避免同时有多个 goroutine 并发修改数据。有的时候，我们也会使用 TryLock。这样一来，当某个 goroutine 想要更改配置数据时，如果发现已经有 goroutine 在更改了，其他的 goroutine 调用 TryLock，返回了 false，这个 goroutine 就会放弃更改。

很多语言（比如 Java）都为锁提供了 TryLock 的方法，但是，Go 官方issue 6123有一个讨论（后来一些 issue 中也提到过），标准库的 Mutex 不会添加 TryLock 方法。虽然通过 Go 的 Channel 我们也可以实现 TryLock 的功能，但是基于 Channel 的实现我们会放在 Channel 那一讲中去介绍，这一次我们还是基于 Mutex 去实现，毕竟大部分的程序员还是熟悉传统的同步原语，而且传统的同步原语也不容易出错。所以这节课，还是希望带你掌握基于 Mutex 实现的方法。

那怎么实现一个扩展 TryLock 方法的 Mutex 呢？我们直接来看代码。

```go
// 复制Mutex定义的常量
const (
    mutexLocked = 1 << iota // 加锁标识位置
    mutexWoken              // 唤醒标识位置
    mutexStarving           // 锁饥饿标识位置
    mutexWaiterShift = iota // 标识waiter的起始bit位置
)

// 扩展一个Mutex结构
type Mutex struct {
    sync.Mutex
}

// 尝试获取锁
func (m *Mutex) TryLock() bool {
    // 如果能成功抢到锁
    if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
        return true
    }

    // 如果处于唤醒、加锁或者饥饿状态，这次请求就不参与竞争了，返回false
    old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
    if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
        return false
    }

    // 尝试在竞争的状态下请求锁
    new := old | mutexLocked
    return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}
```




