package benchmark

import (
	"sync"
	"time"
)

type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (l *RWLock) Write() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.count++
	time.Sleep(COST)
}

func (l *RWLock) Read() {
	l.mu.RLock()
	defer l.mu.RUnlock()
	time.Sleep(COST)
	_ = l.Write
}
