package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	// int类型指针
	aPtr := &a
	// 不支持指针运算
	// aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*") //初始化零值是“”
	t.Log(len(s))
	// 是"",空字符，而非 nil
	if s == "" {
		t.Log("s is null")
	}
}
