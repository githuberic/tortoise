package main

import (
	"log"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	// 返回匿名函数
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

//!+main
func bigSlowOperation() {
	// 延迟执行的函数在return语句之后执行
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

//!-main

func main() {
	bigSlowOperation()
}

