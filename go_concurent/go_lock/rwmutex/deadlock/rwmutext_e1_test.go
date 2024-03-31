package deadlock

import (
	"fmt"
	"sync"
	"testing"
)

func verifyE1() {
	var mu sync.RWMutex

	// writer,稍微等待，然后制造一个调用Lock的场景
	go func() {
		//time.Sleep(200 * time.Millisecond)
		mu.Lock()
		fmt.Println("Lock")
		//time.Sleep(100 * time.Millisecond)
		mu.Unlock()
		fmt.Println("Unlock")
	}()

	go func() {
		factorial(&mu, 10) // 计算10的阶乘, 10!
	}()

	select {}
}

func TestVerifyE1(t *testing.T)  {
	verifyE1()
}