package main

import "fmt"

var count int

func main() {
	for i := 0; i < 2; i++ {
		go func() {
			for i := 1000000; i > 0; i-- {
				count ++
			}
			fmt.Println(count)
		}()
	}

	fmt.Scanf("\n")  //等待子线程全部结束
}
