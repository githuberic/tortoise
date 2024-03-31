package benchmark

import (
	"sync"
	"time"
)

type Lock struct {
	count int
	mu    sync.Mutex
}

func (l *Lock) Write() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.count++
	time.Sleep(COST)
}

func (l *Lock) Read() {
	l.mu.Lock()
	defer l.mu.Unlock()
	time.Sleep(COST)
	_ = l.Write
}
