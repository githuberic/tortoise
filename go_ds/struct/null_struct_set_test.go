package _struct

import (
	"fmt"
	"testing"
)

type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func TestVerifyNSS(t *testing.T) {
	s := make(Set)
	s.Add("lgq")
	s.Add("Tom")
	fmt.Println(s.Has("lgq"))
	fmt.Println(s.Has("Jack"))
}
