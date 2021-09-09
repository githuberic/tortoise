package slice

import (
	"fmt"
	"testing"
	"unsafe"
)

func Change(arr []int) {
	arr = append(arr, 6)
	fmt.Printf("%v %d %d %p\n", arr, len(arr), cap(arr), &arr)
	Printf(arr)
}

/**
slice实质是一个结构体，其作为参数传递时形参实质复制了实参整个结构体的内容，其实就是值传递。
形参分配有一份内存空间，存放和实参相同的内容，从运行结果可以看出形参的内存地址和实参是不同的。

因为形参中底层数组指针和实参相同，所以当做修改操作时会同步修改到实参中，
但是当使用append函数添加元素时，append函数返回的slice会覆盖修改到形参的内存空间中，
和实参无关，所以在main函数中实参不变。可以在上面代码中看到函数中形参已变但实参未变。
 */
func TestVerify(t *testing.T) {
	var arrA = make([]int, 0)
	fmt.Printf("%v %d %d %p\n", arrA, len(arrA), cap(arrA), &arrA)
	Printf(arrA)
	Change(arrA)
	fmt.Printf("%v %d %d %p\n", arrA, len(arrA), cap(arrA), &arrA)
	Printf(arrA)
}


func ChangeV2(arr []int) {
	arr = append(arr, 6)
	fmt.Printf("%v %d %d %p\n", arr, len(arr), cap(arr), &arr)
}
/**
其实实参中tmp[3]已经变为4，但是实参和形参内存空间中len和cap是独立的，
形参中len修改为了4但实参中len仍然为3，所以实参中未增加元素。
 */
func TestV2Verify(t *testing.T) {
	var arr = make([]int, 0, 5)
	arr = append(arr, 1, 2, 3)
	fmt.Printf("%v %d %d %p\n", arr, len(arr), cap(arr), &arr)
	ChangeV2(arr)
	fmt.Printf("%v %d %d %p\n", arr, len(arr), cap(arr), &arr)
}


func ChangeV3(arr []int) {
	arr = append(arr, 6)
	fmt.Printf("%v %d %d %p\n", arr, len(arr), cap(arr), arr)
}
func TestV3Verify(t *testing.T) {
	tmp := make([]int, 0, 5)
	tmp = append(tmp, 1, 2, 3, 4, 5)
	fmt.Printf("%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)
	ChangeV3(tmp[:3])
	fmt.Printf("%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)
}

func ChangeV4(arr []int) {
	arr = append(arr, 4)
	fmt.Printf("%v %d %d %p\n", arr, len(arr), cap(arr), arr)
}
func TestV4Verify(t *testing.T) {
	tmp := make([]int, 0, 5)
	tmp = append(tmp, 1, 2, 3)
	fmt.Printf("%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)

	ChangeV4(tmp)

	// 用实参tmp[2]的地址往后移一个元素地址长度，得到tmp[3]的地址输出，可以看到变为了4。
	p := unsafe.Pointer(&tmp[2])
	q := uintptr(p) + 8
	n := (*int)(unsafe.Pointer(q))
	fmt.Println(*n)
	fmt.Printf("%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)
}

func ChangeV5(arr *[]int) {
	*arr = append(*arr, 4)
	fmt.Printf("%p\n", arr)
	fmt.Printf("%v %d %d %p\n", *arr, len(*arr), cap(*arr), *arr)
}
func TestV5Verify(t *testing.T) {
	tmp := make([]int, 0, 5)
	tmp = append(tmp, 1, 2, 3)
	fmt.Printf("%p\n", &tmp)
	fmt.Printf("%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)
	ChangeV5(&tmp)
	fmt.Printf("%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)
}

/**
当传指针时，对函数中slice的任何修改其实都是对主函数中slice的修改；当传引用，
即slice本身时，对函数中slice使用append时的修改实际是对形参新分配内存空间的修改而实参不变，
但当直接修改slice中值时能同步修改到实参中。
 */

func Printf(arr []int)  {
	fmt.Printf("Value=>%v," +
		"Length=%d," +
		"Capacity=%d," +
		"Pointer=%p," +
		"Address=%v\n",
		arr,
		len(arr),
		cap(arr),
		&arr,
		unsafe.Pointer(&arr))
}