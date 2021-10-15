package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	// 长度不同的数组，不能比较
	//	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)
	// 类型相同、长度相同
	t.Log(a == d)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

const (
	Male = iota
	Female
	Unknown
)

func TestBitClear(t *testing.T) {
	a := 7 //0111
	t.Log("a=,b=,c=", Readable, Writable, Executable)

	t.Log(Unknown)

	// &^ 按位清零运算符
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
