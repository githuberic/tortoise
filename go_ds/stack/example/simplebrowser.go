package example

import (
	"fmt"
	"tortoise/go_ds/stack"
)

type Browser struct {
	forwardStack stack.Stack
	backStack    stack.Stack
}

func NewBrowser() *Browser {
	return &Browser{
		forwardStack: stack.NewArrayStack(),
		backStack:    stack.NewLinkedListStack(),
	}
}

func (this *Browser) CanForward() bool {
	if this.forwardStack.IsEmpty() {
		return false
	}
	return true
}

func (this *Browser) CanBack() bool {
	if this.backStack.IsEmpty() {
		return false
	}
	return true
}

func (this *Browser) Open(addr string) {
	fmt.Printf("Open new addr %+v\n", addr)
	this.forwardStack.Flush()
}

func (this *Browser) PushBack(addr string) {
	this.backStack.Push(addr)
}

func (this *Browser) Forward() {
	if this.forwardStack.IsEmpty() {
		return
	}
	top := this.forwardStack.Pop()
	this.backStack.Push(top)
	fmt.Printf("forward to %+v\n", top)
}

func (this *Browser) Back() {
	if this.backStack.IsEmpty() {
		return
	}
	top := this.backStack.Pop()
	this.forwardStack.Push(top)
	fmt.Printf("back to %+v\n", top)
}
