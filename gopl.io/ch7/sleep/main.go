package main

import (
	"flag"
	"fmt"
	"time"
)

//!+sleep
// flag.Duration 函数创建了一个time.Duration类型的标志变量，并允许用户用一种友好方式来指定时长
var period = flag.Duration("period", 1*time.Second, "sleep period")

/**
run：./main -period 500ms
 */
func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
