package e2

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 递归调用计算阶乘
func factorial(m *sync.RWMutex, n int) int {
	if n < 1 { // 阶乘退出条件
		return 0
	}

	fmt.Println("RLock")
	m.RLock()
	defer func() {
		fmt.Println("RUnlock")
		m.RUnlock()
	}()

	time.Sleep(100 * time.Millisecond)
	return factorial(m, n-1) * n // 递归调用
}

func verify() {

	var mu sync.RWMutex

	// writer,稍微等待，然后制造一个调用Lock的场景
	go func(mu sync.RWMutex) {
		time.Sleep(200 * time.Millisecond)
		mu.Lock()
		fmt.Println("Lock")
		time.Sleep(100 * time.Millisecond)
		mu.Unlock()
		fmt.Println("Unlock")
	}(mu)

	var mu2 sync.RWMutex
	go func(mu sync.RWMutex) {
		factorial(&mu, 10) // 计算10的阶乘, 10!
	}(mu2)

	select {}
}

func TestVerify(t *testing.T)  {
	verify()
}