package go_pointer

import (
	"fmt"
	"testing"
)

func TestPointerE5(t *testing.T) {
	a := 10
	// 取变量a的地址，将指针保存到b中
	b := &a
	fmt.Printf("a=%d, ptr=%p\n", a, &a)
	fmt.Printf("b=%p, typ=e%T\n", b, b)
	fmt.Println(&b)

	// 总结： 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
	// 指针取值（根据指针去内存取值）
	c := *b
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}
