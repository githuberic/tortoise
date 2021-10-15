package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
	data string
}

var singleInstance *Singleton

var once sync.Once

func GetSingletonObj() *Singleton {
	// Only one times
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T)  {
	var wg sync.WaitGroup

	for i :=0; i < 5; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("address =%X\n",unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
