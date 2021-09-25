package multiple_value

import (
	"fmt"
	"testing"
)

func swap(x, y string) (string, string) {
	return y, x
}

func TestVerify(t *testing.T) {
	a, b := swap("A", "B")
	fmt.Println(a, b)
}
