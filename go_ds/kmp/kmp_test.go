package kmp

import (
	"fmt"
	"testing"
)

func TestKMP(t *testing.T)  {
	s := "abc abcdab abcdabcdabde"
	pattern := "bcdabd"
	fmt.Println(findByKMP(s, pattern)) //16

	s = "aabbbbaaabbababbabbbabaaabb"
	pattern = "abab"
	fmt.Println(findByKMP(s, pattern)) //11

	s = "aabbbbaaabbababbabbbabaaabb"
	pattern = "ababacd"
	fmt.Println(findByKMP(s, pattern)) //-1

	s = "hello"
	pattern = "ll"
	fmt.Println(findByKMP(s, pattern)) //2
}
