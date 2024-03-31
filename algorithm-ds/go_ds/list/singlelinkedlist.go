package list

import "fmt"

/*
单链表基本操作
*/
type Node struct {
	Next  *Node
	Value interface{}
}

type SingleLinkedList struct {
	Head   *Node
	Length uint
}

func NewListNode(v interface{}) *Node {
	return &Node{nil, v}
}

func (this *Node) GetNext() *Node {
	return this.Next
}

func (this *Node) GetValue() interface{} {
	return this.Value
}

func NewLinkedList() *SingleLinkedList {
	return &SingleLinkedList{NewListNode(0), 0}
}

// InsertAfter /**
func (s *SingleLinkedList) InsertAfter(p *Node, v interface{}) bool {
	if p == nil {
		return false
	}

	// 构造新节点
	newNode := NewListNode(v)
	// 插入位置(p)node的next节点
	oldNext := p.Next
	// 插入位置(p).next = 新构造的节点
	p.Next = newNode
	// 新构造的节点.next = 插入位置(p)node的next节点
	newNode.Next = oldNext

	s.Length++

	return true
}

// InsertBefore 在某个节点前面插入节点 /**
func (s *SingleLinkedList) InsertBefore(p *Node, v interface{}) bool {
	if nil == p || p == s.Head {
		return false
	}

	cur := s.Head.Next
	pre := s.Head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}

	if nil == cur {
		return false
	}

	newNode := NewListNode(v)
	pre.Next = newNode
	newNode.Next = cur

	s.Length++

	return true
}

// InsertToHead 在链表头部插入节点 /**
func (s *SingleLinkedList) InsertToHead(v interface{}) bool {
	return s.InsertAfter(s.Head, v)
}

// InsertToTail 在链表尾部插入节点 /**
func (s *SingleLinkedList) InsertToTail(v interface{}) bool {
	cur := s.Head
	for nil != cur.Next {
		cur = cur.Next
	}
	return s.InsertAfter(cur, v)
}

// FindByIndex 通过索引查找节点 /**
func (s *SingleLinkedList) FindByIndex(index uint) *Node {
	if index >= s.Length {
		return nil
	}

	cur := s.Head.Next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.Next
	}

	return cur
}

// DeleteNode 删除传入的节点 /**
func (s *SingleLinkedList) DeleteNode(p *Node) bool {
	if nil == p {
		return false
	}

	cur := s.Head.Next
	pre := s.Head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}

	if nil == cur {
		return false
	}

	pre.Next = p.Next
	p = nil
	s.Length--

	return true
}

// Print /**
func (s *SingleLinkedList) Print() {
	cur := s.Head.Next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.Next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
