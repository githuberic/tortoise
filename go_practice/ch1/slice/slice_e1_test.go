package slice

import (
	"fmt"
	"testing"
	"unsafe"
)

func ChangeValue(arr ...int) {
	arr[1] = 10
}
func TestVerifyE1(t *testing.T) {
	var arr = []int{1, 2, 3}
	t.Log("Before change,", "value=", arr, ",address=", unsafe.Pointer(&arr))
	ChangeValue(arr...)
	t.Log("After change,", "value=", arr, ",address=", unsafe.Pointer(&arr))

	a, b, c := 1, 2, 3
	t.Log("Before change,", b, ",address=", unsafe.Pointer(&b))
	ChangeValue(a, b, c)
	t.Log("After change,", b, ",address=", unsafe.Pointer(&b))
}

func appendSlice(arr []int) {
	arr = append(arr, 44)
}
func appendSliceV2(arr []int) []int {
	arr = append(arr, 444)
	return arr
}

// 指针传参
func appendSliceV3(arr *[]int) {
	*arr = append(*arr, 4444)
}
func TestVerifyV2(t *testing.T) {
	var arr = []int{1}
	arr = append(arr, 2, 3, 4, 5)

	appendSlice(arr)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(arr), cap(arr), arr)

	a2 := appendSliceV2(arr)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(a2), cap(a2), a2)

	appendSliceV3(&arr)
	fmt.Printf("len: %d cap:%d data:%+v\n", len(arr), cap(arr), arr)
}


func Change(arr []int) {
	arr = append(arr, 6)
	Printf(arr)
}
/**
slice实质是一个结构体，其作为参数传递时形参实质复制了实参整个结构体的内容，其实就是值传递。形参分配有一份内存空间，存放和实参相同的内容，从运行结果可以看出形参的内存地址和实参是不同的。

因为形参中底层数组指针和实参相同，所以当做修改操作时会同步修改到实参中，

但是当使用append函数添加元素时，append函数返回的slice会覆盖修改到形参的内存空间中，和实参无关，所以在main函数中实参不变。可以在上面代码中看到函数中形参已变但实参未变。
*/
func TestVerify(t *testing.T) {
	var arrA = make([]int, 0)
	Printf(arrA)
	Change(arrA)
	Printf(arrA)
}

func ChangeV2(arr []int) {
	arr = append(arr, 6)
	Printf(arr)
}

/**
其实实参中tmp[3]已经变为4，但是实参和形参内存空间中len和cap是独立的，
形参中len修改为了4但实参中len仍然为3，所以实参中未增加元素。
*/
func TestV2Verify(t *testing.T) {
	var arr = make([]int, 0, 5)
	arr = append(arr, 1, 2, 3)
	Printf(arr)
	ChangeV2(arr)
	Printf(arr)
}

func ChangeV4(arr []int) {
	arr = append(arr, 4)
	Printf(arr)
}
func TestV4Verify(t *testing.T) {
	var arr = make([]int, 0, 5)
	arr = append(arr, 1, 2, 3)
	Printf(arr)

	ChangeV4(arr)

	// 用实参tmp[2]的地址往后移一个元素地址长度，得到tmp[3]的地址输出，可以看到变为了4。
	p := unsafe.Pointer(&arr[2])
	q := uintptr(p) + 8
	n := (*int)(unsafe.Pointer(q))
	fmt.Println(*n)
	Printf(arr)
}

func ChangeV5(arr *[]int) {
	*arr = append(*arr, 4)
	Printf(*arr)
}
func TestV5Verify(t *testing.T) {
	var arr = make([]int, 0, 5)
	arr = append(arr, 1, 2, 3)
	Printf(arr)
	ChangeV5(&arr)
	Printf(arr)
}

/**
当传指针时，对函数中slice的任何修改其实都是对主函数中slice的修改；当传引用，
即slice本身时，对函数中slice使用append时的修改实际是对形参新分配内存空间的修改而实参不变，
但当直接修改slice中值时能同步修改到实参中。
*/
func Printf(arr []int) {
	fmt.Printf("Value=>%v,"+
		"Len=%d,"+
		"Capacity=%d,"+
		"Pointer=%p,"+
		"Address=%v\n",
		arr,
		len(arr),
		cap(arr),
		&arr,
		unsafe.Pointer(&arr))
}
