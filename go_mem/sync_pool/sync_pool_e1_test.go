package sync_pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var bytePool = sync.Pool{
	New: func() interface{} {
		var b = make([]byte, 1024)
		return &b
	},
}


func TestVerifyE1(t *testing.T) {
	a := time.Now().Unix()
	fmt.Println("a", a)
	for i := 0; i < 100000; i++ {
		obj := make([]byte, 1024)
		_ = obj
		fmt.Println("i", i)
	}
	b := time.Now().Unix()
	fmt.Println("b", b)
	fmt.Println("Without pool", b-a, "s")
}

func TestVerify(t *testing.T) {
	a := time.Now().Unix()
	for i := 0; i < 1000000000; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
	b := time.Now().Unix()
	for i := 0; i < 1000000000; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
	c := time.Now().Unix()
	fmt.Println("Without pool", b-a, "s")
	fmt.Println("With pool", c-b, "s")
}
