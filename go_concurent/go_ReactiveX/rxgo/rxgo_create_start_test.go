package rxgo

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func Supplier1(ctx context.Context) rxgo.Item {
	return rxgo.Of(1)
}

func Supplier2(ctx context.Context) rxgo.Item {
	return rxgo.Of(2)
}

func Supplier3(ctx context.Context) rxgo.Item {
	return rxgo.Of(3)
}

func TestStart(t *testing.T)  {
	observable := rxgo.Start([]rxgo.Supplier{Supplier1, Supplier2, Supplier3})
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

