package basic

import (
	"fmt"
	"testing"
)

//Read from close channel
func Test115(t *testing.T) {
	ch := make(chan bool)
	close(ch)

	v, ok := <-ch
	if ok {
		fmt.Printf("Read from close channel: v = %v, ok = %v\n", v, ok)
	}
}
