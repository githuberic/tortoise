package go_error

import (
	"errors"
	"fmt"
	"testing"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New 构造一个使用给定的错误信息的基本error 值。
		// 返回错误值为 nil 代表没有错误。
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg   int
	param string
}

/**
通过实现 Error 方法来自定义 error 类型是可以的。
*/
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.param)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}
	}
	return arg + 3, nil
}

func TestVerify(t *testing.T) {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.param)
	}
}
