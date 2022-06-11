## os/signal包实现对信号的处理
golang中对信号的处理主要使用os/signal包中的两个方法：
- 一个是notify方法用来监听收到的信号；
- 一个是 stop方法用来取消监听。

notify方法：
```go
func Notify(c chan<- os.Signal, sig ...os.Signal)
```
第一个参数表示接收信号的channel, 第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号。

stop方法：
```go
func Stop(c chan<- os.Signal)　　
```