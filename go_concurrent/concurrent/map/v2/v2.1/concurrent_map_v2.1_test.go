package v2_1

import (
	"log"
	"sync"
	"testing"
)

type ConcurrentMap struct {
	sync.RWMutex
	Map map[int]int
}

func (m *ConcurrentMap) Get(index int) (int, bool) {
	m.RLock()
	defer m.RUnlock()
	value, ok := m.Map[index]
	return value, ok
}

func (m *ConcurrentMap) Set(index int, value int) {
	m.Lock()
	defer m.Unlock()
	m.Map[index] = value
}

var mMap *ConcurrentMap

func TestVerify(t *testing.T) {
	mMap = &ConcurrentMap{
		Map: make(map[int]int),
	}

	for i := 0; i < 1000; i++ {
		go func() {
			mMap.Set(i, i)
		}()

		go func() {
			value, ok := mMap.Get(i)
			if ok {
				log.Print(value)
			}
		}()
	}
}
