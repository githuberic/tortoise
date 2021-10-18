package go_tricky

import "testing"

// recover捕获的是祖父级调用时的异常，直接调用时无效：
func TestRecoverDeferV1(t *testing.T) {
	recover()
	panic(1)
}

// 直接defer调用也是无效：
func TestRecoverDeferV2(t *testing.T) {
	defer recover()
	panic(1)
}

// defer调用时多层嵌套依然无效：
func TestRecoverDeferV3(t *testing.T) {
	defer func() {
		func() { recover() }()
	}()
	panic(1)
}

func TestRecoverDeferV4(t *testing.T) {
	defer func() {
		recover()
	}()
	panic(1)
}

