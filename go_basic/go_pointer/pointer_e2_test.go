package go_pointer

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type user struct {
	name string
	age  int
}

func TestVerify(t *testing.T) {
	var u = new(user)
	u.name = "lgq"
	u.age = 18
	fmt.Println(*u)

	// 通过内存偏移的方式，定位到我们需要操作的字段，然后改变他们的值
	// 第一个修改 user 的 name 值的时候，因为 name 是第一个字段，所以不用偏移，
	// 我们获取 user 的指针，然后通过 unsafe.Pointer 转为 *string 进行赋值操作即可。
	var pName = (*string)(unsafe.Pointer(u))
	*pName = "刘**"

	// 第二个修改 user 的 age 值的时候，因为 age 不是第一个字段，所以我们需要内存偏移，内存偏移牵涉到的计算只能通过 uintptr，
	// 所我们要先把 user 的指针地址转为 uintptr，然后我们再通过 unsafe.Offsetof(u.age)
	// 获取需要偏移的值，进行地址运算(+)偏移即可。
	// 现在偏移后，地址已经是 user 的 age 字段了，如果要给它赋值，我们需要把 uintptr 转为 *int 才可以。
	// 所以我们通过把 uintptr 转为 unsafe.Pointer，再转为 *int 就可以操作了。
	var pAge = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge = 19

	fmt.Println(*u)
}

func TestVerifyV3(t *testing.T)  {
	var v1 = uint(12)
	var v2 = int(13)

	fmt.Println(reflect.TypeOf(v1)) // uint
	fmt.Println(reflect.TypeOf(v2)) // int

	fmt.Println(reflect.TypeOf(&v1)) // *uint
	fmt.Println(reflect.TypeOf(&v2)) // *int

	var p = &v1
	p = (*uint)(unsafe.Pointer(&v2)) // 使用unsafe.Pointer进行类型的转换
	*p = 18
	fmt.Println(reflect.TypeOf(p)) // *unit
	fmt.Println(*p)
}