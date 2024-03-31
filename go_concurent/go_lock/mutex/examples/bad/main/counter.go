package main

import "sync"

type Counter struct {
	sync.Mutex
	Count int
}

func f1(c Counter) {
	c.Lock()
	defer c.Unlock()
	println(" f1()")
}

func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	f1(c)
}
