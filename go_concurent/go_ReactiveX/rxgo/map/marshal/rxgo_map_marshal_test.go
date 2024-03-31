package marshal

import (
	"encoding/json"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"testing"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestMarshal(t *testing.T)  {
	observable := rxgo.Just(
		User{
			Name: "dj",
			Age:  18,
		},
		User{
			Name: "jw",
			Age:  20,
		},
	)()

	observable = observable.Marshal(json.Marshal)

	for item := range observable.Observe() {
		// Marshal操作返回的是[]byte类型，我们需要进行类型转换之后再输出。
		fmt.Println(string(item.V.([]byte)))
	}
}