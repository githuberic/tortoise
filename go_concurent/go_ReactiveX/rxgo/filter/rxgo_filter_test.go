package filter

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func TestFilter(t *testing.T){
	observable := rxgo.Range(1, 10)

	observable = observable.Filter(func(i interface{}) bool {
		return i.(int)%2 == 0
	})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func TestElementAt(t *testing.T)  {
	observable := rxgo.Range(1,10).ElementAt(2)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func TestDistinct(t *testing.T){
	observable := rxgo.Just(1, 2, 2, 3, 3, 4, 4)().
		Distinct(func(_ context.Context, i interface{}) (interface{}, error) {
			return i, nil
		})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func TestSkip(t *testing.T)  {
	// Skip可以跳过前若干个数据。
	observable := rxgo.Just(1,2,3,4,5)().Skip(2)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func TestTake(t *testing.T)  {
	// Take只取前若干个数据。
	observable := rxgo.Just(1, 2, 3, 4, 5)().Take(2)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}