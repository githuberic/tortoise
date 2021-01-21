package fifty_newcomer

import (
	"fmt"
	"testing"
)

func TestVerifyV2(t *testing.T) {
	var arrX = [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(arrX)

	fmt.Println(arrX) //prints [1 2 3] (not ok if you need [7 2 3])
}

func TestVerifyV22(t *testing.T) {
	var arrX = []int{1, 2, 3}

	// 更新原始数组数据，请使用数组指针类型。
	func(arr *[]int) {
		(*arr)[0] = 7
		fmt.Println(*arr) //prints [7 2 3]
	}(&arrX)

	fmt.Println(arrX) //prints [7 2 3]
}

func TestVerifyV23(t *testing.T) {
	var arrX = []int{1, 2, 3}

	// 函数获得了切片变量的副本，它仍然引用原始数据。
	func(arr []int){
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(arrX)
	fmt.Println(arrX) //prints [7 2 3]
}