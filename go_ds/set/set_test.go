package set

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"testing"
)

func TestSet(t *testing.T) {
	s1 := mapset.NewSet(1, 2, 3, 4)
	fmt.Println("s1 contains 3: ", s1.Contains(3))
	fmt.Println("s1 contains 5: ", s1.Contains(5))

	// interface 参数，可以传递任意类型
	s1.Add("poloxue")
	fmt.Println("s1 contains poloxue: ", s1.Contains("poloxue"))
	s1.Remove(3)
	fmt.Println("s1 contains 3: ", s1.Contains(3))

	s2 := mapset.NewSet(1, 3, 4, 5)
	// 并集
	fmt.Println(s1.Union(s2))
}
