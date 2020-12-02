package main

import "fmt"

func main() {
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr[0])
	fmt.Println(arr[len(arr)-1])
	for i, v := range arr {
		fmt.Printf("%d,%d\n", i, v)
	}

	for _, v := range arr {
		fmt.Printf("%d\n", v)
	}

	arrV1 := [...]int{1, 2, 3}
	fmt.Printf("%T\n",arrV1)
}
