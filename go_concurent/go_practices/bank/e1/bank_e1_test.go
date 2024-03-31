package e1_test

import (
	"fmt"
	"go_concurent/go_practices/bank/e1"
	"testing"
)

func TestVerify(t *testing.T) {
	done := make(chan struct{})

	go func() {
		e1.Deposit(200)
		fmt.Println("Balance=", e1.Balance())
		done <- struct{}{}
	}()

	go func() {
		e1.Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	if got, want := e1.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
