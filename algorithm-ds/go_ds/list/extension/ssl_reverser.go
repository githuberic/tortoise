package extension

import (
	"algorithm-ds/go_ds/list"
)

type SSLReverser struct {
	list.SingleLinkedList
}

func NewSSLReverser(ssl list.SingleLinkedList) *SSLReverser {
	return &SSLReverser{ssl}
}

func (s *SSLReverser) Reverse() {
	if s.Head == nil || s.Head.Next == nil {
		return
	}

	var pre *list.Node = nil
	var pCur = s.Head.Next

	for pCur != nil {
		tmp := pCur.Next
		pCur.Next = pre
		pre = pCur
		pCur = tmp
	}

	s.Head.Next = pre
}
