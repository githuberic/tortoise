package main

import (
	"github.com/pkg/profile"
	"hash"
	"hash/fnv"
)

var in []byte
var h hash.Hash64

func hashIt() uint64 {
	h.Write(in)
	out := h.Sum64()
	return out
}
func mainFunc() {
	in = []byte("lgq-lgq-lgq-lgq-lgq-lgq-lgq-lgq")
	h = fnv.New64a()
	for i := 0; i < 2000; i++ {
		in = append(in, 'a')
		hashIt()
		h.Reset()
	}
}
func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	mainFunc()
}
