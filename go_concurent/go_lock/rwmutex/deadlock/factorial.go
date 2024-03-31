package deadlock

import (
	"fmt"
	"sync"
)

// 递归调用计算阶乘
func factorial(m *sync.RWMutex, n int) int {
	if n < 1 {
		// 阶乘退出条件
		return 0
	}

	fmt.Println("RLock")
	m.RLock()
	defer func() {
		fmt.Println("RUnlock")
		m.RUnlock()
	}()

	//time.Sleep(100 * time.Millisecond)
	return factorial(m, n-1) * n // 递归调用
}
