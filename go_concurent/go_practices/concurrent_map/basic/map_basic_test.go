package basic

import (
	"fmt"
	"testing"
)

func TestB(t *testing.T) {
	var m = map[string]string{
		"a": "aa",
		"b": "bb",
	}

	for k, v := range m {
		println(k, v)
	}

	if v, ok := m["c"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key Not Found")
	}
}
