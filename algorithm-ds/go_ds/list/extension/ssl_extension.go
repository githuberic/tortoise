package extension

import (
	"algorithm-ds/go_ds/list"
)

type SSLExtension struct {
	list.SingleLinkedList
}

func NewSSLExtension(ssl list.SingleLinkedList) *SSLExtension {
	return &SSLExtension{ssl}
}

// Reverse 单链表反转 /**
func (ssl *SSLExtension) Reverse() {
	if nil == ssl.Head || nil == ssl.Head.Next || nil == ssl.Head.Next.Next {
		return
	}

	// pre 前驱节点
	var pre *list.Node = nil
	// 当前需要逆序的节点
	cur := ssl.Head.Next
	for nil != cur {
		// 将next设置为curr.next
		tmp := cur.Next
		// 将curr.next设置为curr.prev
		cur.Next = pre
		// 将prev和curr的节点依次向后移动一位
		pre = cur
		cur = tmp
	}

	ssl.Head.Next = pre
}

// HasCycle 判断单链表是否有环 /**
func (ssl *SSLExtension) HasCycle() bool {
	if nil != ssl.Head {
		slow := ssl.Head
		fast := ssl.Head
		for nil != fast && nil != fast.Next {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

// MergeSortedList 两个有序单链表合并 /**
func MergeSortedList(l1, l2 *SSLExtension) *SSLExtension {
	if nil == l1 || nil == l1.Head || nil == l1.Head.Next {
		return l2
	}
	if nil == l2 || nil == l2.Head || nil == l2.Head.Next {
		return l1
	}

	l := &SSLExtension{*list.NewLinkedList()}
	cur := l.Head
	curl1 := l1.Head.Next
	curl2 := l2.Head.Next
	for nil != curl1 && nil != curl2 {
		if curl1.Value.(int) > curl2.Value.(int) {
			cur.Next = curl2
			curl2 = curl2.Next
		} else {
			cur.Next = curl1
			curl1 = curl1.Next
		}
		cur = cur.Next
	}

	if nil != curl1 {
		cur.Next = curl1
	} else if nil != curl2 {
		cur.Next = curl2
	}

	return l
}

// DeleteBottomN 删除倒数第N个节点 /**
func (ssl *SSLExtension) DeleteBottomN(n int) {
	if n <= 0 || nil == ssl.Head || nil == ssl.Head.Next {
		return
	}

	fast := ssl.Head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.Next
	}

	if nil == fast {
		return
	}

	slow := ssl.Head
	for nil != fast.Next {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
}

// FindMiddleNode 获取中间节点 /**
func (ssl *SSLExtension) FindMiddleNode() *list.Node {
	if nil == ssl.Head || nil == ssl.Head.Next {
		return nil
	}

	if nil == ssl.Head.Next.Next {
		return ssl.Head.Next
	}

	slow, fast := ssl.Head, ssl.Head
	for nil != fast && nil != fast.Next {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
