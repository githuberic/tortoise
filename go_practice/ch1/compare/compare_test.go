package compare_test

import (
	"fmt"
	"reflect"
	"testing"
)

// 深度比较
/**
复制一个对象时，这个对象可以是内建数据类型、数组、结构体、Map……在复制结构体的时候
如果我们需要比较两个结构体中的数据是否相同，就要使用深度比较，而不只是简单地做浅度比较。
 */
func TestCompare(t *testing.T) {
	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1==m2", reflect.DeepEqual(m1, m2))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1==s2", reflect.DeepEqual(s1, s2))
}
