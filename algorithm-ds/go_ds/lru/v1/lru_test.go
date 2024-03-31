package v1

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	lru := NewLru(5)
	lru.Add("aa", "lgq")
	lru.Add("bb", "lyy")
	lru.Add("cc", "sy")
	lru.Add("dd", "ldx")
	lru.Add("ee", "temp")

	fmt.Println(lru.Get("aa"))
	lru.Remove("ee")

	nodes := lru.GetAll()
	for k, v := range nodes {
		fmt.Println("k,v=", k, v)
	}
}
