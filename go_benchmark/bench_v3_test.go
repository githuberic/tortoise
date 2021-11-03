package go_benchmark

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

/**
直接使用运算符
golang 里面的字符串都是不可变的，每次运算都会产生一个新的字符串，所以会产生很多临时的无用的字符串，
不仅没有用，还会给 gc 带来额外的负担，所以性能比较差
*/
func BenchmarkAddStringWithOperator(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = hello + "," + world
	}
}

/**
fmt.Sprintf()
内部使用 []byte 实现，不像直接运算符这种会产生很多临时的字符串，但是内部的逻辑比较复杂，
有很多额外的判断，还用到了 interface，所以性能也不是很好
*/
func BenchmarkAddStringWithSprintf(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s,%s", hello, world)
	}
}

/**
join会先根据字符串数组的内容，计算出一个拼接之后的长度，然后申请对应大小的内存，
一个一个字符串填入，在已有一个数组的情况下，这种效率会很高，但是本来没有，去构造这个数据的代价也不小
*/
func BenchmarkAddStringWithJoin(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{hello, world}, ",")
	}
}

/**
buffer.WriteString()
这个比较理想，可以当成可变字符使用，对内存的增长也有优化，如果能预估字符串的长度，还可以用 buffer.Grow() 接口来设置 capacity
*/
func BenchmarkAddStringWithBuffer(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < 1000; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(hello)
		buffer.WriteString(",")
		buffer.WriteString(world)
		_ = buffer.String()
	}
}

func BenchmarkAddStringWithAppend(b *testing.B) {
	hello := "hello"
	world := "world"
	for i := 0; i < 1000; i++ {
		_ = string(append([]byte(hello), world...))
	}
}

/**
在已有字符串数组的场合，使用 strings.Join() 能有比较好的性能
在一些性能要求较高的场合，尽量使用 buffer.WriteString() 以获得更好的性能
性能要求不太高的场合，直接使用运算符，代码更简短清晰，能获得比较好的可读性
如果需要拼接的不仅仅是字符串，还有数字之类的其他需求的话，可以考虑 fmt.Sprintf()
*/
