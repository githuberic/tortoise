package rxgo

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
	"time"
)

func TestCreateE1(t *testing.T) {
	observable := rxgo.Create(
		[]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
			next <- rxgo.Of(1)
			next <- rxgo.Of(2)
			next <- rxgo.Of(3)
			next <- rxgo.Error(errors.New("unknown"))
			next <- rxgo.Of(4)
			next <- rxgo.Of(5)
		}})

	ch := observable.Observe()
	for item := range ch {
		if item.Error() {
			fmt.Println("error:", item.E)
		} else {
			fmt.Println(item.V)
		}
	}
}

func TestCreateE2(t *testing.T) {
	observable := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
		next <- rxgo.Error(errors.New("unknown"))
	}, func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(4)
		next <- rxgo.Of(5)
	}})

	ch := observable.Observe()
	for item := range ch {
		if item.Error() {
			fmt.Println("error:", item.E)
		} else {
			fmt.Println(item.V)
		}
	}
}

func TestFromChannel(t *testing.T) {
	ch := make(chan rxgo.Item)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func TestInterval(t *testing.T) {
	observable := rxgo.Interval(rxgo.WithDuration(5 * time.Second))
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	var count int
	for range ticker.C {
		fmt.Println(count)
		count++
	}
}


// Range可以生成一个范围内的数字：
func TestRange(t *testing.T) {
	observable := rxgo.Range(0, 3)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

func TestRepeat(t *testing.T)  {
	observable := rxgo.Just(1, 2, 3)().Repeat(
		3, rxgo.WithDuration(1*time.Second),
	)

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
