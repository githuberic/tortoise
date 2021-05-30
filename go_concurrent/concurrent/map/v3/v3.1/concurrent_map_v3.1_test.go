package v3_1

import (
	"log"
	"testing"
)

type ConcurrentMap struct {
	Map    map[int]int
	ch     chan func()
	tokens chan chan *int
}

func NewConcurrentMap() *ConcurrentMap {
	m := &ConcurrentMap{
		Map:    make(map[int]int),
		ch:     make(chan func()),
		tokens: make(chan chan *int, 128),
	}

	for i := 0; i < cap(m.tokens); i++ {
		m.tokens <- make(chan *int)
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

func (m *ConcurrentMap) find(index int) *int {
	// 每次查询都要创建一个信道
	c := <-m.tokens
	m.ch <- func() {
		res, ok := m.Map[index]
		if !ok {
			c <- nil
			//data = res
		} else {
			inf := res
			c <- &inf
		}
	}
	inf := <-c
	// 回收信道
	m.tokens <- c

	return inf
}

func TestVerify(t *testing.T) {
	mMap := NewConcurrentMap()

	for i := 0; i < 1000; i++ {
		go func() {
			mMap.add(i, i)
		}()

		go func() {
			value := mMap.find(i)
			if value != nil && *value > 0 {
				log.Print(*value)
			}
		}()
		//go mMap.find(i)
	}
}

// https://blog.csdn.net/weixin_34248023/article/details/92487730?utm_medium=distribute.pc_relevant.none-task-blog-baidujs_title-0&spm=1001.2101.3001.4242
