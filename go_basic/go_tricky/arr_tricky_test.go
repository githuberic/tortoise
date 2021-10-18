package go_tricky

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T)  {
	a := [3]int{1,2,3}

	func(arr [3]int){
		arr[0]=7
		fmt.Println(arr)
	}(a)

	fmt.Println(a)
}
