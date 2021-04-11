package e3

import (
	"fmt"
	"sync"
	"testing"
)

func foo(l *sync.RWMutex) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l *sync.RWMutex) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func TestVerify(t *testing.T) {
	l := &sync.RWMutex{}
	foo(l)
}
