package main

import (
	"github.com/pkg/profile"
	"hash"
	"hash/fnv"
)

var h hash.Hash64
var in [2000]byte

func hashIt(i int) uint64 {
	h.Write(in[:i])
	out := h.Sum64()
	return out
}
func mainFunc() {
	h = fnv.New64a()
	for i := 0; i < 2000; i++ {
		in[i] = 'a'
		hashIt(i)
		h.Reset()
	}
}
func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	mainFunc()
}
