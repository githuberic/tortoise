package null_struct

import (
	"fmt"
	"testing"
)

/**
使用空结构体 struct{} 可以节省内存，一般作为占位符使用，表明这里并不需要一个值。
比如使用 map 表示集合时，只关注 key，value 可以使用 struct{} 作为占位符。
如果使用其他类型作为占位符，例如 int，bool，不仅浪费了内存，而且容易引起歧义。
*/
type Set map[string]struct{}

func TestVerify(t *testing.T) {

	var set = make(Set)

	for _, item := range []string{"A", "A", "B", "C"} {
		set[item] = struct{}{}
	}

	fmt.Println(len(set))

	if _, ok := set["A"]; ok {
		fmt.Println("A exists")
	}
}
