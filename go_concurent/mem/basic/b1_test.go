package basic

import "testing"

var a string
var done bool

func setup()  {
	a = "hello world"
	done = true
}

func TestB1(t *testing.T)  {
	go setup()

	for !done{

	}

	println(a)
}


