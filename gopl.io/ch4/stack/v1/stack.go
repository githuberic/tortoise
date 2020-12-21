package main

import "fmt"

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}
func (s stack) Pop() (stack, int) {
	l := len(s)
	if l == 0 {
		return s, 0
	}
	return s[:l-1], s[l-1]
}

func main() {
	s := make(stack, 0)
	s = s.Push(10)
	s = s.Push(11)
	s = s.Push(12)
	fmt.Println(s)
	s, v := s.Pop()
	s, v = s.Pop()
	fmt.Println(s, v)
	s, v = s.Pop()
	fmt.Println(s, v)
}
