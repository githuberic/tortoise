package v2

import "fmt"

type Node struct {
	val  interface{}
	next *Node
}

type QueueLink struct {
	head   *Node
	tail   *Node
	length int
}

func NewQueueLink() *QueueLink {
	return &QueueLink{nil, nil, 0}
}

func (this *QueueLink) EnQueue(v interface{}) bool {
	node := &Node{v, nil}
	if this.tail == nil {
		this.head = node
	} else {
		this.tail.next = node

	}
	this.length++
	this.tail = node

	return true
}

func (this *QueueLink) DeQueue() interface{} {
	if this.head == nil {
		return nil
	}
	v := this.head.val
	this.head = this.head.next
	this.length--
	return v
}

func (this *QueueLink) String() string {
	if this.head == nil {
		return "empty queue"
	}
	result := "head<-"
	for cur := this.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result
}
