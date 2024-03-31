package _map

import (
	"strconv"
	"sync"
	"testing"
)

type MutexMap struct {
	sync.Mutex
	Map map[string]interface{}
}

func (m *MutexMap) readMap(key string) (interface{}, bool) {
	m.Lock()
	defer m.Unlock()
	value, ok := m.Map[key]
	return value, ok
}

func (m *MutexMap) writeMap(index string, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.Map[index] = value
}

func TestVerifyMutex(t *testing.T) {
	var mMap = &MutexMap{
		Map: make(map[string]interface{}),
	}

	for i := 0; i < 1000; i++ {
		key := strconv.Itoa(i)
		go func() {
			mMap.writeMap(key, i)
		}()

		go mMap.readMap(key)
	}
}
