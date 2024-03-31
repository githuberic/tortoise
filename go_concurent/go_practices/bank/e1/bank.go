package e1

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int

	for {
		select {
		case amount := <-deposits:
			balance += amount
			fmt.Printf("balance add,balance=%d\n", balance)
		case balances <- balance:
			fmt.Printf("balance no change,balance=%d\n", balance)
		}
	}
}

func init() {
	fmt.Printf("start init>>>\n")
	go teller()
}
