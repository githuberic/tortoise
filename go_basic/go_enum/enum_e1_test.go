package go_enum

import (
	"fmt"
	"testing"
)

/**
通常使用常量(const) 来表示枚举值。
*/
type StuType int32

const (
	Type1 StuType = iota
	Type2
	Type3
	Type4
)

func TestVerify(t *testing.T) {
	fmt.Println(Type1, Type2, Type3, Type4)
}
