package _map

import (
	"log"
	"sync"
	"testing"
)

func TestVerify(T *testing.T) {
	var m sync.Map

	for i := 1; i < 1000; i++ {
		go func() {
			m.Store(i, i)
		}()

		go func() {
			// loop
			m.Range(func(key, value interface{}) bool {
				k := key.(int)
				v := value.(int)
				log.Println("key=", k, ",value=", v)
				return true
			})
		}()
	}
}
