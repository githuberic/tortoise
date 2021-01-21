package fifty_newcomer

import (
	"fmt"
	"testing"
)

func TestVerifyV1(t *testing.T)  {
	/**
	「nil」标识符可以用作接口，函数，指针，映射，切片和通道的「零值」。
	如果不指定变量类型，则编译器将无法编译代码，因为它无法猜测类型。
	 */
	var x interface{} = nil
	_ = x

	var arr []int
	arr = append(arr,1)
	fmt.Println(arr)
}
