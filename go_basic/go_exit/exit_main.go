package main

import (
	"fmt"
	"time"
)

func tFunc() {
	go func() {
		fmt.Println("father start")
		go func() {
			fmt.Println("son start")
			time.Sleep(time.Second)
			fmt.Println("son exit")
		}()
		fmt.Println("father exit")
	}()
}

func main() {
	fmt.Println("main start")
	tFunc()
	time.Sleep(time.Second * 2)
	go func() {
		fmt.Println("t1 start")
		fmt.Println("t1 exit")
	}()

	fmt.Println("main exit")
}
