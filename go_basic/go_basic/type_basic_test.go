package go_basic

import (
	"fmt"
	"testing"
)

func TestTypeBasic(t *testing.T) {
	n := 121331
	v := fmt.Sprintf("%v", n)
	fmt.Println(v)
}
