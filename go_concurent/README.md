# go_concurent
golang 相关的并发安全practices

## go并发编程风格
* go routine Channel
goroutine和通道channel，他们支持通信顺序进程(Communicating Sequential process CSP),CSP 是一种并发模式，在不同的执行体
(goroutine)之间传递值；
在go里面，每一个并发执行的活动称为goroutine；程序启动时，只有一个goroutine来调用main函数，称之为主goroutine; 新的goroute
可以通过go语句进行创建
f() 调用函数f();等待它返回
go f() 新建一个调用f()的goroutine 不用等待

### 实例1
主goroutine执行计算斐波那契数列，在其执行期间，我们开启一个goroutine，显示一个字符串(spinner)来指示程序依然在运行
