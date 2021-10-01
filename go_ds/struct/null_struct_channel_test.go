package _struct

import (
	"fmt"
	"testing"
)

func worker(ch chan struct{})  {
	<-ch
	fmt.Println("do something>>>")
	close(ch)
}

func TestVerifyNSC(t *testing.T)  {
	ch := make(chan struct{})
	go worker(ch)
	ch<- struct{}{}
	fmt.Println("done>>>")
}
