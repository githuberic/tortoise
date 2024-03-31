# Channel

### 概述

Channel 是 Go 语言内建的 first-class 类型，也是 Go 语言与众不同的特性之一。

Go 语言的 Channel 设计精巧简单，以至于也有人用其它语言编写了类似 Go 风格的 Channel 库，比如docker/libchan、tylertreat/chan，

但是并不像 Go 语言一样把 Channel 内置到了语言规范中。

### Channel 的发展

CSP 是 Communicating Sequential Process 的简称，中文直译为通信顺序进程，或者叫做交换信息的循序进程，是用来描述并发系统中进行交互的一种模式。

<b>CSP 允许使用进程组件来描述系统，它们独立运行，并且只通过消息传递的方式通信。</b>

#### Channel 的应用场景

<b>：“执行业务处理的 goroutine 不要通过共享内存的方式通信，而是要通过 Channel 通信的方式分享数据。”</b>

- “communicate by sharing memory”和“share memory by communicating”是两种不同的并发处理模式。“communicate by sharing memory”是传统的并发编程处理方式，就是指，共享的数据需要用锁进行保护，goroutine 需要获取到锁，才能并发访问数据。
- “share memory by communicating”则是类似于 CSP 模型的方式，通过通信的方式，一个 goroutine 可以把数据的“所有权”交给另外一个 goroutine（虽然 Go 中没有“所有权”的概念，但是从逻辑上说，你可以把它理解为是所有权的转移）。

Channel 的应用场景分为五种类型
- 数据交流：当作并发的 buffer 或者 queue，解决生产者 - 消费者问题。多个 goroutine 可以并发当作生产者（Producer）和消费者（Consumer）。
- 数据传递：一个 goroutine 将数据交给另一个 goroutine，相当于把数据的拥有权 (引用) 托付出去。
- 信号通知：一个 goroutine 可以将信号 (closing、closed、data ready 等) 传递给另一个或者另一组 goroutine 。
- 任务编排：可以让一组 goroutine 按照一定的顺序并发或者串行的执行，这就是编排的功能。
- 锁：利用 Channel 也可以实现互斥锁的机制。

##### Channel 基本用法

```go
chan string          // 可以发送接收string
chan<- struct{}      // 只能发送struct{}
<-chan int           // 只能从chan接收int
```
# Channel 使用场景

[golang常见使用场景][https://github.com/Shitaibin/shitaibin.github.io/blob/hexo_resource/source/_posts/golang-channel-all-usage.md]reference-style
