package go_mem

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestVerify(t *testing.T) {
	str := "qcon广州"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))

	for index, value := range str {
		fmt.Printf("%d,%c\n", index, value)
	}

	// 字符串拼接 内部使用[]byte实现涉及到类型转换，可以拼接其他类型，性能一般，少量非字符串类型拼接时
	s := fmt.Sprint("Qconf","广州",2019)
	fmt.Println(s)

	// 只能拼接字符串数组 不灵活,已存在字符串数组时
	s1 := []string{"Qconf","广州","2019"}
	fmt.Println(strings.Join(s1,""))

	// 拼接字符串、字符和unicode,底层使用[]byte,涉及到string和[]byte之间转换,少量字符串拼接时
	var b bytes.Buffer
	b.WriteString("Qconf")
	b.WriteString("广州")
	b.WriteString("2019")
	s2 := b.String()
	fmt.Println(s2)

	// 拼接字符串、字符和unicode 使用unsafe.Pointer优化了string和[]byte之间 的转换,大量字符串拼接
	var b2 strings.Builder
	b2.WriteString("Qconf")
	b2.WriteString("广州")
	b2.WriteString("2019")
	s3 := b.String()
	fmt.Println(s3)
}
