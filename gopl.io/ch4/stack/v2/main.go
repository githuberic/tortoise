package main

import (
	"errors"
	"fmt"
	"sync"
)

type stack struct {
	lock sync.Mutex
	arr  []int
}

func NewStack() *stack {
	return &stack{sync.Mutex{}, make([]int, 10)}
}

func (stk *stack) Push(v int) {
	stk.lock.Lock()
	defer stk.lock.Unlock()

	stk.arr = append(stk.arr, v)
}

func (stk *stack) Pop() (int, error) {
	stk.lock.Lock()
	defer stk.lock.Unlock()

	len := len(stk.arr)
	if len == 0 {
		return -1, errors.New("Empty stack")
	}
	v := stk.arr[len-1]
	stk.arr = stk.arr[:len-1]
	return v, nil
}

func main() {
	fmt.Println(">>>")
	s := NewStack()
	s.Push(10)
	s.Push(11)
	s.Push(12)
	fmt.Println(s.arr)
	vv, _ := s.Pop()
	vv, _ = s.Pop()
	fmt.Println(s.arr, vv)
}
