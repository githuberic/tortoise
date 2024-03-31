package e3

import "sync"

type BankBiz struct {
	lock    sync.Mutex
	balance int
}

func NewBankBiz() *BankBiz {
	return &BankBiz{lock: sync.Mutex{}}
}

func (this *BankBiz) Deposit(amount int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.balance += amount
}

func (this *BankBiz) Balance() int {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.balance
}
