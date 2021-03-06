package bank

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
			fmt.Printf("balance add,balance=%d\n", balance)
		case balances <- balance:
			fmt.Printf("balance no change,balance=%d\n",balance)
		}
	}
}

func init() {
	// teller /goroutine限制balance变量
	go teller() // start the monitor goroutine
}
