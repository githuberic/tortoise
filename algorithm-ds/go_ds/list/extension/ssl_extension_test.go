package extension

import (
	"algorithm-ds/go_ds/list"
	"testing"
)

var ssl *SSLExtension

func init() {
	l := list.NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToHead(i + 1)
	}

	ssl = NewSSLExtension(*l)
}

func TestReverse(t *testing.T) {
	ssl.Print()
	ssl.Reverse()
	ssl.Print()
}

func TestHasCycle(t *testing.T) {
	t.Log(ssl.HasCycle())
	ssl.Head.Next.Next.Next.Next.Next.Next = ssl.Head.Next.Next.Next
	t.Log(ssl.HasCycle())
}

func TestMergeSortedList(t *testing.T) {
	l1 := list.NewLinkedList()
	for i := 0; i < 5; i++ {
		l1.InsertToHead(i + 1)
	}
	ssl1 := NewSSLExtension(*l1)

	l2 := list.NewLinkedList()
	for i := 5; i < 10; i++ {
		l2.InsertToHead(i + 1)
	}
	ssl2 := NewSSLExtension(*l2)

	MergeSortedList(ssl1,ssl2).Print()
}


func TestDeleteBottomN(t *testing.T) {
	ssl.Print()
	ssl.DeleteBottomN(3)
	ssl.Print()
}

func TestFindMiddleNode(t *testing.T) {
	ssl.DeleteBottomN(1)
	ssl.Print()
	t.Log(ssl.FindMiddleNode())
}
