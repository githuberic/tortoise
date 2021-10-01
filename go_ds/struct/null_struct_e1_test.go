package _struct

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestVerify(t *testing.T) {
	// 空结构体 struct{} 实例不占据任何的内存空间。
	fmt.Println(unsafe.Sizeof(struct{}{}))
}
