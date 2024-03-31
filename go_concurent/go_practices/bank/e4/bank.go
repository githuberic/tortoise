package e4

import "sync"

type BankBiz struct {
	lock    sync.RWMutex
	balance int
}

func NewBankBiz() *BankBiz {
	return &BankBiz{lock: sync.RWMutex{}}
}

func (this *BankBiz) Deposit(amount int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.balance += amount
}

func (this *BankBiz) Balance() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.balance
}
