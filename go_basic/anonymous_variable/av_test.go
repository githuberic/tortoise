package test

import (
	"fmt"
	"testing"
)

func GetData(x, y int) (int, int) {
	return x, x + y
}

func TestVerifyAV(t *testing.T) {
	a, _ := GetData(1, 2)
	_, sum := GetData(1, 2)
	fmt.Println(a, sum)
}
