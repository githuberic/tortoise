package _map

import (
	"log"
	"strconv"
	"sync"
	"testing"
)

type RWMutexMap struct {
	sync.RWMutex
	Map map[string]interface{}
}

func (m *RWMutexMap) Get(index string) (interface{}, bool) {
	m.RLock()
	defer m.RUnlock()
	value, ok := m.Map[index]
	return value, ok
}

func (m *RWMutexMap) Set(index string, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.Map[index] = value
}

func TestVerifyRWMutex(t *testing.T) {
	var mMap = &RWMutexMap{
		Map: make(map[string]interface{}),
	}

	for i := 1; i < 1000; i++ {
		key := strconv.Itoa(i)
		go func() {
			mMap.Set(key, i)
		}()

		go func() {
			value, ok := mMap.Get(key)
			if ok {
				log.Print(value)
			}
		}()
	}
}
