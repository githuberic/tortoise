package go_pointer

import (
	"fmt"
	"testing"
	"unsafe"
)

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

type StringHeader struct {
	Data uintptr
	Len  int
}

func Bytes2String(b []byte) string {
	var sliceHeader = (*SliceHeader)(unsafe.Pointer(&b))

	var sh = StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))
}

func String2Byte(s string) []byte {
	var stringHeader = (*StringHeader)(unsafe.Pointer(&s))

	var bh = SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}

func TestVerifyVp(t *testing.T)  {
	var s = "Hello World"
	var b = String2Byte(s)
	fmt.Println(b)
	s = Bytes2String(b)
	fmt.Println(s)
}
