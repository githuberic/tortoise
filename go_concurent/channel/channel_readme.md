# Channel
通道 goroutine之间的连接，可以让一个goroutine 发送特定值到另外一个goroutine的通信机制

每一个通道是一个具体类型的导管 叫做通道元素，e，g，int类型的通道 chan int

使用内置的make函数来创建一个通道
ch := make(chan int)  // ch的类型 是chan int

通道主要有两种通信操作 
- 发送send<br/>
  发送语句 ch < -x
- 接收receive，两个操作都可以使用<-操作符号<br/> 
  x = <-ch 赋值语句中接收表达式<br/>
  <-ch 接收语句，丢弃结果

- close()关闭操作：<br/>
  设置一个标识位指示值当前已经发送完毕，这个通道后面值了；关闭后发送操作将导致宕机；

## 信道创建
make函数创建无缓冲(unbuffered)通道
ch = make(chan int) // 无缓冲通道
ch = make(chan int,0) // 无缓冲通道
ch = make(chan int,3) // 容量为3缓冲通道

使用无缓冲(unbuffered)通道进行通信会导致发送和接收goroutine同步化（同步通道）;





