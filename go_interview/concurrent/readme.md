### golang 的channel是怎么实现的

golang的channel是个结构体，里面大概包含了三大部分：
- a. 指向内容的环形缓存区，及其相关游标;
- b. 读取和写入的排队goroutine链表;
- c. 锁;

任何操作前都需要获得锁， 当写满或者读空的时候，就将当前goroutine加入到recvq或者sendq中， 并出让cpu(gopark)。

### 什么是 Channel，为什么它可以做到线程安全？

Channel 是Go中的一个核心类型，可以把它想象成一个可以用于传递数据的通道，不过是用在并发编程中的， Channel 也可以理解是一个先进先出的队列，通过管道进行通信。

Golang 的 Channel， 发送一个数据到 Channel 和从 Channel 接收一个数据都是原子性的。

Go 的设计思想就是, 不要通过共享内存来通信，而是通过通信来共享内存，前者就是传统的加锁，后者就是 Channel。

也就是说，设计 Channel 的主要目的就是在多任务间传递数据的，本身就是线程安全的。