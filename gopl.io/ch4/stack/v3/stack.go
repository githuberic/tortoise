package stack

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (stk *Stack) Len() int {
	return stk.length
}

// View the top item on the stack
func (stk *Stack) Peek() interface{} {
	if stk.length == 0 {
		return nil
	}
	return stk.top.value
}

// Pop the top item of the stack and return it
func (stk *Stack) Pop() interface{} {
	if stk.length == 0 {
		return nil
	}
	n := stk.top
	stk.top = n.prev
	stk.length--
	return n.value
}

// Push a value onto the top of the stack
func (stk *Stack) Push(value interface{}) {
	n := &node{value, stk.top}
	stk.top = n
	stk.length++
}
