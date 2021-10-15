package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i :=0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i) // 值传递，i值被复制了一份, 和其他go routine没有竞争关系
	}
	time.Sleep(time.Millisecond * 50)
}
