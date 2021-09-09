# 普通指针、unsafe.Pointer、uintptr的区别和联系

## unsafe
unsafe 是不安全的，所以我们应该尽可能少的使用它，比如内存的操纵，这是绕过 Go 本身设计的安全机制的，不当的操作，可能会破坏一块内存，而且这种问题非常不好定位。

当然必须的时候我们可以使用它，比如底层类型相同的数组之间的转换；比如使用 sync/atomic 包中的一些函数时；还有访问 struct 的私有字段时；该用还是要用，不过一定要慎之又慎。

## golang指针
- *类型:普通指针类型，用于传递对象地址，不能进行指针运算。
- unsafe.Pointer:通用指针类型，用于转换不同类型的指针，不能进行指针运算，不能读取内存存储的值（必须转换到某一类型的普通指针）。unsafe.Pointer 是桥梁，可以让任意类型的指针实现相互转换，也可以将任意类型的指针转换为 uintptr 进行指针运算。
- uintptr:用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收。

### unsafe.Pointer四个规则：
- 任何类型的指针都可以被转化为Pointer
- Pointer可以被转化为任何类型的指针
- uintptr可以被转化为Pointer
- Pointer可以被转化为uintptr

### 指针类型转换
Go 语言在设计的时候，被设计成为一门强类型的静态语言。强类型意味着一旦定义了，它的类型就不能改变了；静态意味着类型检查在运行前就做了。
同时为了安全的考虑，Go 语言是不允许两个指针类型进行转换的。

### unsafe.pointer用于普通指针类型转换
```go
func TestVerifyV2(t *testing.T) {
	var i = 10
	var ip = &i

	var fp *int64 = (*int64)(unsafe.Pointer(ip))
	*fp = *fp * 3

	fmt.Println(i)
	fmt.Println(fp)
	//fmt.Println(&fp)
	fmt.Println(*fp)
}
```
### 指针转为uintptr偏移计算
*我们都知道 T 是不能计算偏移量的，也不能进行计算，但是 uintptr 可以，所以我们可以把指针转为 uintptr 再进行偏移计算，这样我们就可以访问特定的内存了，达到对不同的内存读写的目的。
```go
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
```


