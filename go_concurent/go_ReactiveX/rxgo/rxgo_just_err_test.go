package rxgo

import (
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)


func TestJustError(t *testing.T) {
	observable := rxgo.Just(1, 2, errors.New("unknown"), 3, 4, 5)()
	ch := observable.Observe()
	for item := range ch {
		// rxgo.Item还可以包含错误。所以在使用时，我们应该做一层判断：
		if item.Error() {
			fmt.Println("error:", item.E)
		} else {
			fmt.Println(item.V)
		}
	}
}
