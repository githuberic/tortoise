package extension

import (
	"algorithm-ds/go_ds/list"
	"testing"
)

var sslr *SSLReverser

func init() {
	l := list.NewLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToHead(i + 1)
	}

	sslr = NewSSLReverser(*l)
}

func TestReverser(t *testing.T) {
	sslr.Print()
	sslr.Reverse()
	sslr.Print()
}

