package main

import (
	"fmt"
	"time"
)

func f11() {
	for {
		fmt.Println("call f1...")
	}
}

func f22() {
	fmt.Println("call f2...")
}

func main() {
	//go f22()
	in := make(chan int)

	go func(ch chan int) {
		ch <- 11
	}(in)

	go func(ch chan int) {
		//for {
			fmt.Println(<-ch)
		//}
	}(in)

	time.Sleep(time.Millisecond)

	//close(in)
	/*
		select {
		case <-in:
			//fmt.Println(<-in)
		default:
			fmt.Println("exit")
		}*/
}
