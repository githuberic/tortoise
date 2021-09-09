package go_pointer

import (
	"fmt"
	"testing"
	"unsafe"
)

/*
func TestVerify(t *testing.T) {
	var i = 10
	var ip = &i

	// 会提示 cannot convert ip (type *int) to type *float64，也就是不能进行强制转型
	var fp *float64 = (*float64)(ip)
	fmt.Println(fp)
}
*/

func TestVerifyV2(t *testing.T) {
	var i = 10
	var ip = &i

	var fp *int64 = (*int64)(unsafe.Pointer(ip))
	*fp = *fp * 3

	fmt.Println(i)
	fmt.Println(fp)
	//fmt.Println(&fp)
	fmt.Println(*fp)
}
