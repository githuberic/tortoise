package ring

import (
	"container/ring"
	"fmt"
	"testing"
)

type SumInt struct {
	Value int
}

func (s *SumInt) add(i interface{}) {
	s.Value += i.(int)
}

func TestRingDo(t *testing.T) {
	r := ring.New(10)

	for i := 0; i < 10; i++ {
		r.Value = i
		r = r.Next()
	}

	sum := SumInt{}
	r.Do(sum.add)
	fmt.Println(sum.Value)
}
