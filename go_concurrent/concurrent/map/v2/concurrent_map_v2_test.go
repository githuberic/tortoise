package v2

import (
	"sync"
	"testing"
)

type ConcurrentMap struct {
	sync.Mutex
	Map map[int]int
}

func (m *ConcurrentMap) readMap(index int) (int, bool) {
	m.Lock()
	value, ok := m.Map[index]
	m.Unlock()
	return value, ok
}

func (m *ConcurrentMap) writeMap(index int, value int) {
	m.Lock()
	m.Map[index] = value
	m.Unlock()
}

var mMap *ConcurrentMap

func TestVerify(t *testing.T) {
	mMap = &ConcurrentMap{
		Map: make(map[int]int),
	}

	for i := 0; i < 1000; i++ {
		go func() {
			mMap.writeMap(i, i)
		}()

		go mMap.readMap(i)
	}
}

func readMap(index int) (int, bool) {
	return mMap.readMap(index)
}
