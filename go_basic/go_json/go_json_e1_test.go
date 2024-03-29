package go_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type ResponseV1 struct {
	Page   int
	Fruits []string
}

type ResponseV2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func TestVerify(t *testing.T) {
	// 首先我们看一下将基础数据类型编码为JSON数据
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("lgq")
	fmt.Println(string(strB))

	// 这里是将切片和字典编码为JSON数组或对象
	var slcD = []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	var mapD = map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON包可以自动地编码自定义数据类型。结果将只包括自定义
	// 类型中的可导出成员的值并且默认情况下，这些成员名称都作
	// 为JSON数据的键
	var res1 = &ResponseV1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	resv1B, _ := json.Marshal(res1)
	fmt.Println(string(resv1B))

	var res2 = &ResponseV2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	resv2B, _ := json.Marshal(res2)
	fmt.Println(string(resv2B))

	// 现在我们看看解码JSON数据为Go数值
	var byt = []byte(`{"num":6.13,"strs":["a","b"]}`)

	// 我们需要提供一个变量来存储解码后的JSON数据，这里
	// 的`map[string]interface{}`将以Key-Value的方式
	// 保存解码后的数据，Value可以为任意数据类型
	var dat map[string]interface{}

	// 解码过程，并检测相关可能存在的错误
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 为了使用解码后map里面的数据，我们需要将Value转换为
	// 它们合适的类型，例如我们将这里的num转换为期望的float64
	var num = dat["num"].(float64)
	fmt.Println(num)

	// 访问嵌套的数据需要一些类型转换
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 我们还可以将JSON解码为自定义数据类型，这有个好处是可以
	// 为我们的程序增加额外的类型安全并且不用再在访问数据的时候
	// 进行类型断言
	var str = `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &ResponseV2{}
	json.Unmarshal([]byte(str),&res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 上面的例子中，我们使用bytes和strings来进行原始数据和JSON数据
	// 之间的转换，我们也可以直接将JSON编码的数据流写入`os.Writer`
	// 或者是HTTP请求回复数据。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
