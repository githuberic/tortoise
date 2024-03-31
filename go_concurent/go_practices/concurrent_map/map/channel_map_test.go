package _map

import (
	"log"
	"testing"
)

type ChannelMap struct {
	Map map[int]int
	ch  chan func()
}

func NewChannelMap() *ChannelMap {
	m := &ChannelMap{
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

func (m *ChannelMap) add(index int, value int) {
	m.ch <- func() {
		m.Map[index] = value
	}
}

func (m *ChannelMap) del(index int) {
	m.ch <- func() {
		delete(m.Map, index)
	}
}

func (m *ChannelMap) find(index int) (data int) {
	// 每次查询都要创建一个信道
	m.ch <- func() {
		if res, ok := m.Map[index]; ok {
			data = res
		}
	}
	return
}

func TestVerifyChMap(t *testing.T) {
	mMap := NewChannelMap()

	for i := 0; i < 1000; i++ {
		go func() {
			mMap.add(i, i)
		}()

		go func() {
			value := mMap.find(i)
			if value > 0 {
				log.Printf("Value=%d",value)
			}
		}()
	}
}
