package main

import (
	"fmt"
	"net/url"
)

func main() {
	//!+main
	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"      (first value)
	fmt.Println(m["item"])     // "[1 2]"  (direct map access)

	// m=空map
	m = nil
	// m.Get("item") 不能通过编译,nil的类型没有确定
	fmt.Println(m.Get("item")) // ""
	// 发生crash 更新空的map
	m.Add("item", "3")         // panic: assignment to entry in nil map
	//!-main
}
