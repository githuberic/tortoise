package go_interview

import "testing"

func f1() *int {
	var v = 11
	return &v
}

func TestVerifyEEA(t *testing.T)  {
	var f = f1()
	println(*f)
}
