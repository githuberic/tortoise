package v1_1

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestVerify1(t *testing.T) {
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
