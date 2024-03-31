package e4

import (
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	bz := NewBankBiz()
	var wg sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func(amount int) {
			bz.Deposit(amount)
			wg.Done()
		}(i)
	}
	wg.Wait()

	if got, want := bz.Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
