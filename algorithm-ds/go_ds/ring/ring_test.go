package ring

import (
	"container/ring"
	"testing"
)

func TestRing(t *testing.T) {
	r := ring.New(10)
	for i := 0; i < 10; i++ {
		r.Value = i
		r = r.Next()
	}
}
