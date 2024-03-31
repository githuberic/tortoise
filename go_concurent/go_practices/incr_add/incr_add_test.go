package incr_add

import (
	"fmt"
	"testing"
	"time"
)

var count = 0
var count0 = 0

func TestIncrAdd(t *testing.T) {
	for i := 0; i < 100; i++ {
		go addUnsafe()
	}
	time.Sleep(time.Second * 1)
	fmt.Println(">>>Count0", count0)

	chs := make([]chan int, 100)
	for i := 0; i < 100; i++ {
		chs[i] = make(chan int, 1)
		go addSafe(i, chs[i])
	}

	for _, ch := range chs {
		v := <-ch
		count += v
	}
	fmt.Println(">>>Count", count)
}

func addUnsafe() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum++
	}
	count0 += sum
}

func addSafe(j int, ch chan int) {
	sum := 0
	for i := 0; i < 100; i++ {
		sum++
	}
	ch <- sum
}
