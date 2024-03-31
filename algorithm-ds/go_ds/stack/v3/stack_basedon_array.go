package v3

import "fmt"

type Stack struct {
	//数据
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{
		data: make([]interface{}, 0, 32),
	}
}

func (this *Stack) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *Stack) Push(v interface{}) {
	this.data = append(this.data, v)
}

func (this *Stack) PushList(v []interface{}) {
	this.data = append(this.data, v...)
}

func (this *Stack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	ret := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return ret
}

func (this *Stack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.data[len(this.data)-1]
}

func (this *Stack) Flush() {
	if this.IsEmpty() {
		return
	}
	for i := 0; i < len(this.data)-1; i++ {
		this.data[i] = nil
	}
	this.data = make([]interface{}, 0, 32)
}

func (this *Stack) Print() {
	if this.IsEmpty() {
		fmt.Println("Empty stack")
	} else {
		for i := len(this.data) - 1; i >= 0; i-- {
			fmt.Println(this.data[i])
		}
	}
}
