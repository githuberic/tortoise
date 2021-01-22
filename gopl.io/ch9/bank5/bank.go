package bank

import "sync"

type account struct {
	mu      sync.Mutex // guards balance
	balance int
}

//var acc = new(account)
var acc = &account{}

func Withdraw(amount int) bool {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	deposit(-amount)
	if acc.balance < 0 {
		deposit(amount)
		return false // 余额不足
	}
	return true
}

func Deposit(amount int) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	deposit(amount)
}

func Balance() int {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	return acc.balance
}

func deposit(amount int) {
	acc.balance += amount
}

