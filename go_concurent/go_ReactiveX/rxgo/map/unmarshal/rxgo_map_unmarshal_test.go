package unmarshal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	observable := rxgo.Just(
		`{"name":"dj","age":18}`,
		`{"name":"jw","age":20}`,
	)()

	// Unmarshaller接受[]byte类型的参数，我们在Unmarshal之前加了一个Map用于将string转为[]byte。运行：
	observable = observable.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		return []byte(i.(string)), nil
	}).Unmarshal(json.Unmarshal, func() interface{} {
		return &User{}
	})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
