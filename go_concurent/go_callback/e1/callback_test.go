package e1

import (
	"fmt"
	"testing"
)

type Callback func(x, y int) int

func callbackFunc(x, y int, callback Callback) int {
	return callback(x, y)
}

func Add(x, y int) int {
	return x + y
}

func TestCallback(t *testing.T) {
	x, y := 1, 2
	fmt.Println(callbackFunc(x, y, Add))
}
