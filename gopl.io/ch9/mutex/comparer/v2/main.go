package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	//读写锁
	countGuard sync.RWMutex
)

func read(mapA map[string]string) {
	for {
		countGuard.RLock() //这里定义了一个读锁
		var _ string = mapA["name"]
		count += 1
		countGuard.RUnlock()
	}
}

func write(mapA map[string]string) {
	for {
		countGuard.Lock() //这里定义了一个写锁
		mapA["name"] = "johny"
		count += 1
		time.Sleep(time.Millisecond * 3)
		countGuard.Unlock()
	}
}

func main() {
	var num int = 3
	var mapA map[string]string = map[string]string{"nema": ""}

	for i := 0; i < num; i++ {
		go read(mapA)
	}

	for i := 0; i < num; i++ {
		go write(mapA)
	}

	time.Sleep(time.Second * 3)
	fmt.Printf("最终读写次数：%d\n", count)
}
