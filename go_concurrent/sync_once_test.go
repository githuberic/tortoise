package go_concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestVerifyV2(t *testing.T)  {
	var once sync.Once
	one := func() {
		fmt.Println("just once")
	}

	for i := 0; i < 10; i++ {
		go func(a int) {
			once.Do(one)   // 只是被执行一次
		}(i)
	}
	time.Sleep(time.Millisecond*200)
}
