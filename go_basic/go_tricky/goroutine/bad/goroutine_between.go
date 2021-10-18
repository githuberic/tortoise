package go_bad

var msg string
var done bool

func setup()  {
	msg = "hello world"
	done = true
}

// 因为在不同的Goroutine，main函数中无法保证能打印出hello, world:
func main()  {
	go setup()

	for !done {
	}

	println(msg)
}
