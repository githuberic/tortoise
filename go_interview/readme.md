### Go 语言的局部变量分配在栈上还是堆上？
由编译器决定。
Go 语言编译器会自动决定把一个变量放在栈还是放在堆，编译器会做逃逸分析(escape analysis)，
当发现变量的作用域没有超出函数范围，就可以在栈上，反之则必须分配在堆上。

```go
import "testing"

func f1() *int {
	var v = 11
	return &v
}

func TestVerifyEEA(t *testing.T)  {
	var f = f1()
	println(*f)
}
```
f1() 函数中，如果 v 分配在栈上，f1 函数返回时，&v 就不存在了，但是这段函数是能够正常运行的。
Go 编译器发现 v 的引用脱离了 foo 的作用域，会将其分配在堆上。因此，main 函数中仍能够正常访问该值。

### 2 个 interface 可以比较吗？
Go 语言中，interface 的内部实现包含了 2 个字段，类型 T 和 值 V，interface 可以使用 == 或 != 比较。
2 个 interface 相等有以下 2 种情况
- 两个 interface 均等于 nil（此时 V 和 T 都处于 unset 状态）
- 类型 T 相同，且对应的值 V 相等。

### 两个 nil 可能不相等吗？
接口(interface) 是对非接口值(例如指针，struct等)的封装，内部实现包含 2 个字段，类型 T 和 值 V。
一个接口等于 nil，当且仅当 T 和 V 处于 unset 状态（T=nil，V is unset）。
- 两个接口值比较时，会先比较 T，再比较 V。
- 接口值与非接口值比较时，会先将非接口值尝试转换为接口值，再比较。

```go
import "testing"

func TestVerifyNEQ(t *testing.T) {
	var p *int = nil
	var i interface{} = p

	println(i == p)   // true
	println(p == nil) // true
	println(i == nil) // false
}
```
