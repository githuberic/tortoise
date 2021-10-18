package unsafe_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

type Customer struct {
	Name string
	Age  int
}

func TestUnsafe(t *testing.T) {
	i := 10
	t.Log(unsafe.Pointer(&i))

	// 危险的类型转换
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log(f)

	g := *(*int)(unsafe.Pointer(&i))
	t.Log(g)
}


// The cases is suitable for unsafe
type MyInt int

//合理的类型转换
func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)

	c :=10
	c1 := *(*MyInt)(unsafe.Pointer(&c))
	t.Log(c1)
}

//原子类型操作
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 3; i++ {
				writeDataFn()
				//time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			for i := 0; i < 3; i++ {
				readDataFn()
				//time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
