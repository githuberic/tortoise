## Semaphore内部实现

Go 扩展库中的信号量是使用互斥锁 +List 实现的。互斥锁实现其它字段的保护，而 List 实现了一个等待队列，等待者的通知是通过 Channel 的通知机制实现的。

信号量 Weighted 的数据结构：

```go
type Weighted struct {
    size    int64         // 最大资源数
    cur     int64         // 当前已被使用的资源
    mu      sync.Mutex    // 互斥锁，对字段的保护
    waiters list.List     // 等待队列
}
```

在信号量的几个实现方法里，Acquire 是代码最复杂的一个方法，它不仅仅要监控资源是否可用，而且还要检测 Context 的 Done 是否已关闭。我们来看下它的实现代码。

```go

func (s *Weighted) Acquire(ctx context.Context, n int64) error {
    s.mu.Lock()
        // fast path, 如果有足够的资源，都不考虑ctx.Done的状态，将cur加上n就返回
    if s.size-s.cur >= n && s.waiters.Len() == 0 {
      s.cur += n
      s.mu.Unlock()
      return nil
    }
  
        // 如果是不可能完成的任务，请求的资源数大于能提供的最大的资源数
    if n > s.size {
      s.mu.Unlock()
            // 依赖ctx的状态返回，否则一直等待
      <-ctx.Done()
      return ctx.Err()
    }
  
        // 否则就需要把调用者加入到等待队列中
        // 创建了一个ready chan,以便被通知唤醒
    ready := make(chan struct{})
    w := waiter{n: n, ready: ready}
    elem := s.waiters.PushBack(w)
    s.mu.Unlock()
  

        // 等待
    select {
    case <-ctx.Done(): // context的Done被关闭
      err := ctx.Err()
      s.mu.Lock()
      select {
      case <-ready: // 如果被唤醒了，忽略ctx的状态
        err = nil
      default: 通知waiter
        isFront := s.waiters.Front() == elem
        s.waiters.Remove(elem)
        // 通知其它的waiters,检查是否有足够的资源
        if isFront && s.size > s.cur {
          s.notifyWaiters()
        }
      }
      s.mu.Unlock()
      return err
    case <-ready: // 被唤醒了
      return nil
    }
  }
```

其实，为了提高性能，这个方法中的 fast path 之外的代码，可以抽取成 acquireSlow 方法，以便其它 Acquire 被内联。

Release 方法将当前计数值减去释放的资源数 n，并唤醒等待队列中的调用者，看是否有足够的资源被获取。

```go
func (s *Weighted) Release(n int64) {
    s.mu.Lock()
    s.cur -= n
    if s.cur < 0 {
      s.mu.Unlock()
      panic("semaphore: released more than held")
    }
    s.notifyWaiters()
    s.mu.Unlock()
}
```

notifyWaiters 方法就是逐个检查等待的调用者，如果资源不够，或者是没有等待者了，就返回：

```go
func (s *Weighted) notifyWaiters() {
    for {
      next := s.waiters.Front()
      if next == nil {
        break // No more waiters blocked.
      }
  

      w := next.Value.(waiter)
      if s.size-s.cur < w.n {
        //避免饥饿，这里还是按照先入先出的方式处理
        break
      }

      s.cur += w.n
      s.waiters.Remove(next)
      close(w.ready)
    }
  }
```

notifyWaiters 方法是按照先入先出的方式唤醒调用者。当释放 100 个资源的时候，如果第一个等待者需要 101 个资源，那么，队列中的所有等待者都会继续等待，即使有的等待者只需要 1 个资源。这样做的目的是避免饥饿，否则的话，资源可能总是被那些请求资源数小的调用者获取，这样一来，请求资源数巨大的调用者，就没有机会获得资源了。




