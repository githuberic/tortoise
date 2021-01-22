package fifty_newcomer

import (
	"fmt"
	"testing"
)

// 访问不存在的映射键

// 错误的范例:
func TestVerifyV3(t *testing.T) {
	x := map[string]string{"one":"a","two":"","three":"c"}
	if v := x["two"]; v == "" {
		fmt.Println("no entry")
	}
}

// 正确的写法:
func TestVerifyV31(t *testing.T) {
	x := map[string]string{"one":"a","two":"","three":"c"}
	if _, ok := x["two"]; !ok {
		fmt.Println("no entry")
	}
}