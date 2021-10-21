package slice_test

import (
	"bytes"
	"fmt"
	"testing"
	"unsafe"
)

// 切片可变长,cap不够时，cap < 1024, cap容量成倍增加；cap >= 1024时，按照1.25倍扩容
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		// capacity 每次都是2倍增幅
		s = append(s, i)
		t.Log("length=", len(s), "capacity=", cap(s))
	}
}

func TestSlice1(t *testing.T) {
	//var rd = rand.New(rand.NewSource(time.Now().UnixNano()))
	var foo = make([]int, 3, 5)
	// append()的操作，这个操作会让 a 重新分配内存，这就会导致 a 和 b 不再共享
	foo = append(foo, 11, 22, 3, 4, 5, 56, 57, 34, 68)
	Printf(foo)

	// 截取部分
	bar := foo[1:4]
	Printf(bar)

	// 重新赋值
	bar[1] = 10

	// foo 和 bar 的内存是共享的，所以，foo 和 bar 对数组内容的修改都会影响到对方
	Printf(foo)
	Printf(bar)
}

func TestSlice2(t *testing.T) {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	fmt.Printf("Index=%d\n", sepIndex)

	var dir1 = path[:sepIndex]
	fmt.Printf("Path=>%s,Dir1=>%s,Len=%d,Capacity=%d,Address=%v\n", string(path), string(dir1), len(dir1), cap(dir1), unsafe.Pointer(&dir1))

	var dir2 = path[sepIndex+1:]
	fmt.Printf("Path=>%s,Dir2=>%s,Len=%d,Capacity=%d,Address=%v\n", string(path), string(dir2), len(dir2), cap(dir2), unsafe.Pointer(&dir2))

	dir1 = append(dir1, "suffix"...)
	fmt.Printf("Path=>%s,Dir1=>%s,Len=%d,Capacity=%d,Address=%v\n", string(path), string(dir1), len(dir1), cap(dir1), unsafe.Pointer(&dir1))
	fmt.Printf("Path=>%s,Dir2=>%s,Len=%d,Capacity=%d,Address=%v\n", string(path), string(dir2), len(dir2), cap(dir2), unsafe.Pointer(&dir2))
}

func Printf(arr []int) {
	fmt.Printf("Value=>%v,Len=%d,Capacity=%d,Pointer=%p,Address=%v\n", arr, len(arr), cap(arr), &arr, unsafe.Pointer(&arr))
}
