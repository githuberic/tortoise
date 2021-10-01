package go_interview

import "testing"

func TestVerifyNEQ(t *testing.T) {
	var p *int = nil
	var i interface{} = p

	println(i == p)   // true
	println(p == nil) // true
	println(i == nil) // false
}

/**
将一个 nil 非接口值 p 赋值给接口 i，此时，i 的内部字段为(T=*int, V=nil)，
i 与 p 作比较时，将 p 转换为接口后再比较，因此 i == p，p 与 nil 比较，直接比较值，所以 p == nil。

但是当 i 与 nil 比较时，会将 nil 转换为接口 (T=nil, V=nil)，与i (T=*int, V=nil) 不相等，
因此 i != nil。因此 V 为 nil ，但 T 不为 nil 的接口不等于 nil。
*/
