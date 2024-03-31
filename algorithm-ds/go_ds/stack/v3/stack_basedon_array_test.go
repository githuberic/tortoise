package v3

import "testing"

func TestArrayStack_Push(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	t.Log(s.Pop())
	s.Push(3)
	t.Log(s.Pop())
	t.Log(s.Pop())
	s.Push(4)
	t.Log(s.Pop())
	s.Print()
}

func TestArrayStack_Pop(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()

	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	s.Print()
}
