package string_trim

import (
	"fmt"
	"strings"
	"testing"
)

func TestVerifyV2(t *testing.T) {
	// 原始字符串，前后有空格
	str := "   Hello, world!   \n\r "

	// 使用 TrimSpace 函数删除前后空格
	trimmedStr := strings.TrimSpace(str)

	fmt.Println("原始字符串:", str)
	fmt.Println("修剪后字符串:", trimmedStr)
}
