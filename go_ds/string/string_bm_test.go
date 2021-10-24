package string

import (
	"fmt"
	"testing"
)

func TestStringBM(t *testing.T)  {
	main := "abcacabcbcabcabc"
	pattern := "cabcab"

	fmt.Println(bmSearch(main, pattern))
}
