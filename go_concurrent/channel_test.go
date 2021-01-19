package go_concurrent

import (
	"fmt"
	"testing"
	"time"
)

var chanStr = make(chan string, 520)
var arr1 = []string{"qq", "ww", "ee", "rr", "tt"}

func chanTest1() {
	for _, v := range arr1 {
		chanStr <- v
	}
	close(chanStr)
}

func chanTest2() {
	for {
		str, ok := <-chanStr
		if !ok {
			return
		}
		fmt.Println(str)
	}
}

func TestVerifyV5(t *testing.T)  {
	chanTest1()
	chanTest2()

	time.Sleep(time.Millisecond*200)
}
