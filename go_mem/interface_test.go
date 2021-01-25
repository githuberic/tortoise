package go_mem

import (
	"fmt"
	"strconv"
	"testing"
)

func TestVerifyV3(t *testing.T) {
	var b Binary = 200
	// 从何而来
	s := Stringer(b)
	fmt.Println(s.String())
}

type Stringer interface {
	String() string
}

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}
func (i Binary) Get() uint64 {
	return uint64(i)
}
