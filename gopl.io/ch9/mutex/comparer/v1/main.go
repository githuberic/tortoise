package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count  int
	//互斥锁
	countGuard sync.Mutex
)

func read(mapA map[string]string){
	for {
		countGuard.Lock()
		var _ string = mapA["name"]
		count += 1
		countGuard.Unlock()
	}
}

func write(mapA map[string]string) {
	for {
		countGuard.Lock()
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
