package rxgo

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func TestRxGoE1(t *testing.T)  {
	observable := rxgo.Just(1, 2, 3, 4, 5)()
	ch := observable.Observe()
	for item := range ch {
		fmt.Println(item.V)
	}
}