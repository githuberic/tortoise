package rafaeldias

import (
	"errors"
	"fmt"
	"github.com/rafaeldias/async"
	"testing"
	"time"
)

func TestConcurrent(t *testing.T)  {
	res, e := async.Concurrent(async.Tasks{
		func() int {
			for i := 'a'; i < 'a'+26; i++ {
				fmt.Printf("%c ", i)
			}
			return 0
		},
		func() error {
			time.Sleep(3 * time.Microsecond)
			for i := 0; i < 27; i++ {
				fmt.Printf("%d ", i)
			}
			return errors.New("Error executing concurently")
		},
	})

	if e != nil {
		fmt.Printf("Errors [%s]\n", e.Error()) // output errors separated by space
	}

	fmt.Printf("Result from function 0: %v", res.Index(0))
}