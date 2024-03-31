package string

import (
	"fmt"
	"testing"
)

func TestStringBf(t *testing.T)  {
	main := "abcd227fac"
	pattern := "ac"
	fmt.Println(bfSearch(main, pattern))
}
