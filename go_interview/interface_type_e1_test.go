package go_interview

import (
	"fmt"
	"testing"
)

type T string

func (t *T) f1() {
	fmt.Println("Hello")
}

func TestVerifyITET(t *testing.T)  {
	var t1 T = "lgq"
	t1.f1()
	
	const t2 T = "lgq"
	//t2.f1() // cannot call pointer method on t2
}
