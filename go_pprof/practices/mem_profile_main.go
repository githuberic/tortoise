package main

import (
	"hash/fnv"
)

var in []byte

func hashIt() uint64 {
	h := fnv.New64a()
	h.Write(in)
	out := h.Sum64()
	return out
}
func mainFunc() {
	in = []byte("lgq-lgq-lgq-lgq-lgq-lgq-lgq-lgq")
	for i := 0; i < 2000; i++ {
		in = append(in, 'a')
		hashIt()
	}
}
func main() {
	//defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	mainFunc()
}
