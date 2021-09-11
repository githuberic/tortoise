package go_align

import (
	"fmt"
	"testing"
	"unsafe"
)

type A struct {
	a bool
	b string
	c bool
}
/**
A内存占用：32个字节
bool 1 byte 补7byte
string 8*2bytes（指针8bytes len 8byte 2个int）
bool 1 byte 补7byte
 */

type B struct {
	a bool
	c bool
	b string
}
/**
B内存占用：24个字节
bool 1 byte
bool 1 byte
补 8-1-1=6个字节
string 8*2bytes（指针8bytes len 8byte 2个int）
*/

type C struct {
	a bool
	b int8
	c uint16
	d uint32
	e int64
	f bool
}

type D struct {
	a bool
	b int8
	d uint32
	c uint16
	e int64
	f bool
}

func TestVerifyV1(t *testing.T)  {
	fmt.Printf("Size A: %d\n",unsafe.Sizeof(A{}))
	fmt.Printf("Size A: %d\n",unsafe.Sizeof(B{}))

	var b bool = true
	var i int = 40
	fmt.Println("bool-size",unsafe.Alignof(b))
	fmt.Println("int-size",unsafe.Alignof(i))

	type st struct {
		b bool
		i int
	}
	var tmp = st{}
	fmt.Println("st.b",unsafe.Alignof(tmp.b))
	fmt.Println("st.i",unsafe.Alignof(tmp.i))
	fmt.Println("st",unsafe.Sizeof(tmp))

	var c = C{}
	fmt.Println("c占用的实际内存大小:", unsafe.Sizeof(c), "字节,结构体对齐保证:", unsafe.Alignof(c))
	fmt.Println("a:", unsafe.Sizeof(c.a), "字节,字段对齐保证:", unsafe.Alignof(c.a), ",偏移地址:", unsafe.Offsetof(c.a))
	fmt.Println("b:", unsafe.Sizeof(c.b), "字节,字段对齐保证:", unsafe.Alignof(c.b), ",偏移地址:", unsafe.Offsetof(c.b))
	fmt.Println("c:", unsafe.Sizeof(c.c), "字节,字段对齐保证:", unsafe.Alignof(c.c), ",偏移地址:", unsafe.Offsetof(c.c))
	fmt.Println("d:", unsafe.Sizeof(c.d), "字节,字段对齐保证:", unsafe.Alignof(c.d), ",偏移地址:", unsafe.Offsetof(c.d))
	fmt.Println("e:", unsafe.Sizeof(c.e), "字节,字段对齐保证:", unsafe.Alignof(c.e), ",偏移地址:", unsafe.Offsetof(c.e))
	fmt.Println("f:", unsafe.Sizeof(c.f), "字节,字段对齐保证:", unsafe.Alignof(c.f), ",偏移地址:", unsafe.Offsetof(c.f))
	fmt.Println(uintptr(unsafe.Pointer(&c)))

	var d = D{}
	fmt.Println("d占用的实际内存大小:", unsafe.Sizeof(d), "字节,结构体对齐保证:", unsafe.Alignof(d))
	fmt.Println("a:", unsafe.Sizeof(d.a), "字节,字段对齐保证:", unsafe.Alignof(d.a), ",偏移地址:", unsafe.Offsetof(d.a))
	fmt.Println("b:", unsafe.Sizeof(d.b), "字节,字段对齐保证:", unsafe.Alignof(d.b), ",偏移地址:", unsafe.Offsetof(d.b))
	fmt.Println("c:", unsafe.Sizeof(d.c), "字节,字段对齐保证:", unsafe.Alignof(d.c), ",偏移地址:", unsafe.Offsetof(d.c))
	fmt.Println("d:", unsafe.Sizeof(d.d), "字节,字段对齐保证:", unsafe.Alignof(d.d), ",偏移地址:", unsafe.Offsetof(d.d))
	fmt.Println("e:", unsafe.Sizeof(d.e), "字节,字段对齐保证:", unsafe.Alignof(d.e), ",偏移地址:", unsafe.Offsetof(d.e))
	fmt.Println("f:", unsafe.Sizeof(d.f), "字节,字段对齐保证:", unsafe.Alignof(d.f), ",偏移地址:", unsafe.Offsetof(d.f))
	fmt.Println(uintptr(unsafe.Pointer(&d)))
}