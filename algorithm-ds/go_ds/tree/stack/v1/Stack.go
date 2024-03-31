package v1

import "container/list"

type Stack struct {
	list *list.List
}

// NewStack /**
func NewStack() *Stack {
	return &Stack{list.New()}
}

// Push /**
func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

// Pop /**
func (stack *Stack) Pop() interface{} {
	if e := stack.list.Back(); e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

// Len /**
func (stack *Stack) Len() int {
	return stack.list.Len()
}

// Empty /**
func (stack *Stack) Empty() bool {
	return stack.Len() == 0
}
