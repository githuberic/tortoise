package v3

import (
	"log"
	"testing"
)

type ConcurrentMap struct {
	Map map[int]int
	ch  chan func()
}

func NewConcurrentMap() *ConcurrentMap {
	m := &ConcurrentMap{
		Map: make(map[int]int),
		ch:  make(chan func()),
	}

	go func() {
		for {
			(<-m.ch)()
		}
	}()
	return m
}

func (m *ConcurrentMap) add(index int, value int) {
	m.ch <- func() {
		m.Map[index] = value
	}
}

func (m *ConcurrentMap) del(index int) {
	m.ch <- func() {
		delete(m.Map, index)
	}
}

func (m *ConcurrentMap) find(index int) (data int) {
	// 每次查询都要创建一个信道
	m.ch <- func() {
		if res, ok := m.Map[index]; ok {
			data = res
		}
	}
	return
}

func TestVerify(t *testing.T) {
	mMap := NewConcurrentMap()

	for i := 0; i < 1000; i++ {
		go func() {
			mMap.add(i, i)
		}()

		go func() {
			value := mMap.find(i)
			if value > 0 {
				log.Print(value)
			}
		}()
		//go mMap.find(i)
	}
}
