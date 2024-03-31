package deadlock

import (
	"fmt"
	"sync"
	"testing"
)



func verifyE2() {

	var mu sync.RWMutex

	// writer,稍微等待，然后制造一个调用Lock的场景
	go func(mu *sync.RWMutex) {
		//time.Sleep(200 * time.Millisecond)
		mu.Lock()
		fmt.Println("Lock")
		//time.Sleep(100 * time.Millisecond)
		mu.Unlock()
		fmt.Println("Unlock")
	}(&mu)

	/*
	var mu2 sync.RWMutex
	go func(mu sync.RWMutex) {
		factorial(&mu, 10) // 计算10的阶乘, 10!
	}(mu2)*/

	select {}
}

func TestVerify(t *testing.T)  {
	verifyE2()
}