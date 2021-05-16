package object_pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	var pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	var v = pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestVerifyV2(t *testing.T) {
	var pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return "Starter"
		},
	}

	var v = pool.Get().(string)
	fmt.Println(v)

	pool.Put("hello world")
	var v2 = pool.Get().(string)
	fmt.Println(v2)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 1
		},
	}

	pool.Put(10)
	pool.Put(100)
	pool.Put(1000)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
