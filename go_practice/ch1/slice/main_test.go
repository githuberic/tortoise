package slice_test

import (
	"bytes"
	"fmt"
	"testing"
	"unsafe"
)

func TestSlice1(t *testing.T) {
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	bar := foo[1:4]
	bar[1] = 99
	// foo 和 bar 的内存是共享的，所以，foo 和 bar 对数组内容的修改都会影响到对方
	fmt.Println(foo)
	fmt.Println(bar)
}

func TestSlice11(t *testing.T) {
	arrA := make([]int, 32)
	arrB := arrA[1:16]
	// 对 a 做了一个 append()的操作，这个操作会让 a 重新分配内存，这就会导致 a 和 b 不再共享
	arrA = append(arrA, 1)
	arrA = append(arrA, 2)
	arrA[3] = 42
	fmt.Println(arrA)
	fmt.Println(arrB)
}

func TestSlice2(t *testing.T) {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	fmt.Printf("Index=%d\n", sepIndex)

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	t.Log("length=", len(dir1), "capacity=", cap(dir1))

	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB
	t.Log("length=", len(dir2), "capacity=", cap(dir2))

	dir1 = append(dir1, "suffix"...)
	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	t.Log("length=", len(dir1), "capacity=", cap(dir1))

	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB
	t.Log("length=", len(dir2), "capacity=", cap(dir2))
}

func TestVerify(t *testing.T) {
	arr := []int{1, 2, 3}
	t.Log("Before change,",arr[1],",address=",unsafe.Pointer(&arr))
	changeValue(arr...)
	t.Log("After change,",arr[1],",address=",unsafe.Pointer(&arr))

	a,b,c := 1,2,3
	t.Log("Before change,",b,",address=",unsafe.Pointer(&b))
	changeValue(a,b,c)
	t.Log("After change,",b,",address=",unsafe.Pointer(&b))
}

func changeValue(arr ...int)  {
	arr[1] = 0
	fmt.Println("Change,",arr[1],",address=",unsafe.Pointer(&arr))
}
