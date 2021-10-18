package go_good

var msg string
var done = make(chan bool)

// 解决的办法是用显式同步：
func setup() {
	msg = "hello world"
	done <- true
}

// msg的写入是在channel发送之前，所以能保证打印hello, world
func main() {
	go setup()
	<-done
	println(msg)
}
