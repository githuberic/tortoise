package v1

import (
	"container/list"
	"errors"
	"sync"
)

type node struct {
	key   interface{}
	value interface{}
}

type Lru struct {
	max      int
	dataList *list.List
	dataMap  map[interface{}]*list.Element
	rwlock   sync.RWMutex
}

func NewLru(len int) *Lru {
	return &Lru{
		max:      len,
		dataList: list.New(),
		dataMap:  make(map[interface{}]*list.Element),
	}
}

func (l *Lru) Add(key interface{}, val interface{}) error {
	l.rwlock.Lock()
	defer l.rwlock.Unlock()

	if l.dataList == nil {
		return errors.New("empty data list")
	}

	// 已存在，修改value后移至队首
	if e, ok := l.dataMap[key]; ok {
		e.Value.(*node).value = val
		l.dataList.MoveToFront(e)
		return nil
	}

	// 不存在，在队首插入新结点，在map中存储节点
	ele := l.dataList.PushFront(&node{
		key:   key,
		value: val,
	})
	l.dataMap[key] = ele

	// 如果队列超过最大距离，移除队尾的节点，移除map中的key
	if l.max != 0 && l.dataList.Len() > l.max {
		if e := l.dataList.Back(); e != nil {
			l.dataList.Remove(e)
			delete(l.dataMap, e.Value.(*node).key)
		}
	}

	return nil
}

func (l *Lru) Remove(key interface{}) bool {
	l.rwlock.Lock()
	defer l.rwlock.Unlock()

	if l.dataList == nil {
		return false
	}

	// 如果存在节点，先移至队首，再返回value
	if ele, ok := l.dataMap[key]; ok {
		l.dataList.Remove(ele)
		delete(l.dataMap, ele.Value.(*node).key)
		return true
	}

	return false
}

func (l *Lru) Get(key interface{}) (val interface{}, ok bool) {
	l.rwlock.RLock()
	defer l.rwlock.RUnlock()

	if l.dataList == nil {
		return nil, false
	}

	// 如果存在节点，先移至队首，再返回value
	if ele, ok := l.dataMap[key]; ok {
		l.dataList.MoveToFront(ele)
		return ele.Value.(*node).value, true
	}

	return nil, false
}

func (l *Lru) GetAll() []*node {
	l.rwlock.Lock()
	defer l.rwlock.Unlock()

	var data []*node
	for _, v := range l.dataMap {
		data = append(data, v.Value.(*node))
	}
	return data
}
