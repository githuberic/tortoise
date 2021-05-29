package v1

import "testing"

var mMap map[int]int

func TestVerify(t *testing.T) {
	mMap = make(map[int]int)

	for i := 0; i < 1000; i++ {
		go func() {
			mMap[i] = i
		}()

		go readMap(i)
	}
}

func readMap(index int) int {
	return mMap[index]
}
