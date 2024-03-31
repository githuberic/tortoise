package e2

type BankBiz struct {
	sema    chan struct{}
	balance int
}

func NewBankBiz() *BankBiz {
	// chan=1 相当于lock
	return &BankBiz{sema: make(chan struct{}, 1)}
}

func (this *BankBiz) Deposit(amount int) {
	// acquire token
	this.sema <- struct{}{}
	this.balance += amount
	// release token
	<-this.sema
}

func (this *BankBiz) Balance() int {
	this.sema <- struct{}{}
	b := this.balance
	<-this.sema
	return b
}
