package e2

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	v := new(int)
	modifyFunc(v)
	fmt.Println("Main", v)
}

func modifyFunc(v *int) {
	v = nil
	fmt.Println("modifyFunc:", v)
}
