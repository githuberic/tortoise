package _map

import (
	"fmt"
	"log"
	"testing"
	"time"
)

var mMap map[int]int

func TestVerifyE1(t *testing.T) {
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


func TestVerifyE2(t *testing.T) {
	m := make(map[string]int)

	go func() {
		for i := 0; i < 1000; i++ {
			m[fmt.Sprintf("%d", i)] = i
		}
	}()

	go func() {
		for j := 0; j < 1000; j++ {
			log.Println(m[fmt.Sprintf("%d", j)])
		}
	}()

	time.Sleep(time.Second * 20)
}