package ch2

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday)

	a := 1
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
