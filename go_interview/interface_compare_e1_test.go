package go_interview

import (
	"fmt"
	"testing"
)

type Stu struct {
	Name string
}

type StuInt interface{}

func TestVerify(t *testing.T) {
	// stu1 和 stu2 对应的类型是 *Stu，值是 Stu 结构体的地址，两个地址不同，因此结果为 false。
	var stu1, stu2 StuInt = &Stu{"lgq"}, &Stu{"lgq"}
	// stu3 和 stu4 对应的类型是 Stu，值是 Stu 结构体，且各字段相等，因此结果为 true。
	var stu3, stu4 StuInt = Stu{"lgq"}, Stu{"lgq"}
	fmt.Println(stu1 == stu2)
	fmt.Println(stu3 == stu4)
}
