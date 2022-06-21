# 概述
关于参数传递,Golang文档中

After they are evaluated, the parameters of the call are passed by value to the function and the called function begins execution.

函数调用参数均为值传递，不是指针传递或引用传递。经测试引伸出来，当参数变量为指针或隐式指针类型，参数传递方式也是传值（指针自己的copy）

Slice是最经常使用的数据结构之一,下面以Slice为例,解释Golang的参数传递机制

示例代码

```go
import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	slice := make([]int, 3, 5)
	fmt.Println("before", slice)
	changeSliceMember(slice)
	fmt.Println("alter", slice)
}

func changeSliceMember(slice []int) {
	if len(slice) > 0 {
		slice[0] = 88
	}
}
```

函数执行结果为

befor:[0 0 0]
after:[88 0 0]

解释
从数据结构图中可看出,Slice能够理解成结构体类型,包含底层数组首元素地址、数组len、容量三个字段,slice对象在参数传值过程当中,把三个字段的值传递过去了,实际上changeSliceMember函数内slice在内存中的地址和main中的slice内存地址不同,只是字段值是同样的,而第一个字段Pointer的值就是底层数组首元素地址,所以能够直接改变元素内容指针

<br/>
示例代码二

```go
import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	v := new(int)
	modifyFunc(v)
	fmt.Println("Main", v)
}

func modifyFunc(v *int) {
	v = nil
	fmt.Println("modifyFunc:", v)
}
```
能够看出,即便传值为指针,仍未改变变量value在main中的值,由于modifyFunc中value的值为指针,和main中的value值同样,可是俩对象自己是两个对象