package slice_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unsafe"
)

func TestSlice1(t *testing.T) {
	var rd = rand.New(rand.NewSource(time.Now().UnixNano()))
	var foo = make([]int, 5, 10)
	foo[3] = rd.Int()
	foo[4] = rd.Int()
	Printf(foo)

	bar := foo[1:4]
	bar[1] = rd.Int()

	// foo 和 bar 的内存是共享的，所以，foo 和 bar 对数组内容的修改都会影响到对方
	Printf(foo)
	Printf(bar)
}



func TestSlice11(t *testing.T) {
	arrA := make([]int, 10)
	Printf(arrA)

	// 对 a 做了一个 append()的操作，这个操作会让 a 重新分配内存，这就会导致 a 和 b 不再共享
	arrA = append(arrA,1,2,3,4,5,56,57,34,68)
	Printf(arrA)

	arrB := arrA[1:5]
	arrA[3] = 42
	Printf(arrA)

	Printf(arrB)
}

func TestSlice2(t *testing.T) {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	fmt.Printf("Index=%d\n", sepIndex)

	var dir1 = path[:sepIndex]
	fmt.Printf("Path=>%s,Dir1=>%s,Length=%d,Capacity=%d,Address=%v\n",string(path),string(dir1),len(dir1),cap(dir1),unsafe.Pointer(&dir1))

	var dir2 = path[sepIndex+1:]
	fmt.Printf("Path=>%s,Dir2=>%s,Length=%d,Capacity=%d,Address=%v\n",string(path),string(dir2),len(dir2),cap(dir2),unsafe.Pointer(&dir2))

	dir1 = append(dir1, "suffix"...)
	fmt.Printf("Path=>%s,Dir1=>%s,Length=%d,Capacity=%d,Address=%v\n",string(path),string(dir1),len(dir1),cap(dir1),unsafe.Pointer(&dir1))
	fmt.Printf("Path=>%s,Dir2=>%s,Length=%d,Capacity=%d,Address=%v\n",string(path),string(dir2),len(dir2),cap(dir2),unsafe.Pointer(&dir2))
}

func TestVerify(t *testing.T) {
	arr := []int{1, 2, 3}
	t.Log("Before change,", arr, ",address=", unsafe.Pointer(&arr))
	changeValue(arr...)
	t.Log("After change,", arr, ",address=", unsafe.Pointer(&arr))

	a, b, c := 1, 2, 3
	t.Log("Before change,", b, ",address=", unsafe.Pointer(&b))
	changeValue(a, b, c)
	t.Log("After change,", b, ",address=", unsafe.Pointer(&b))
}

func changeValue(arr ...int) {
	arr[1] = 10
	//fmt.Println("Change,", arr[1], ",address=", unsafe.Pointer(&arr))
}

func Printf(arr []int)  {
	fmt.Printf("Value=>%v,Length=%d,Capacity=%d,Address=%v\n",arr, len(arr), cap(arr),unsafe.Pointer(&arr))
}
