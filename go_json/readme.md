##GO_JSON
Go语言在解析来源为非强类型语言时比如PHP等序列化的JSON时，经常遇到一些问题诸如字段类型变化导致无法正常解析的情况，导致服务不稳定。

- 就是挖掘Golang解析json的绝大部分能力
- 比较优雅的解决解析json时存在的各种问题
- 深入一下Golang解析json的过程

### Golang解析JSON之Tag篇

1. 一个结构体正常序列化过后是什么样的呢？

```go
package e1

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Product 商品信息
type Product struct {
	Name      string
	ProductID int64
	Number    int
	Price     float64
	IsOnSale  bool
}

func TestVerifyJSON(t *testing.T) {
	var p = &Product{}
	p.Name = "Phone13"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 8499.00
	p.ProductID = 1001
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	fmt.Println(string(data))
}
```

2. 何为Tag，tag就是标签，给结构体的每个字段打上一个标签，标签冒号前是类型，后面是标签名

```go
package e2

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Product 商品信息
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"-"` // 表示不进行序列化
	Number    int     `json:"number"`
	Price     float64 `json:"price"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

func TestVerifyJSON(t *testing.T) {
	var p = &Product{}
	p.Name = "Phone13"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 8499.00
	p.ProductID = 1001
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	fmt.Println(string(data))
}

// 结果 {"name":"Phone13","number":10000,"price":8499,"is_on_sale":"true"}
```

3. omitempty，tag里面加上omitempy，可以在序列化的时候忽略0值或者空值

```go
package e3

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Product 商品信息
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,omitempty"`
	Number    int     `json:"number"`
	Price     float64 `json:"price"`
	IsOnSale  bool    `json:"is_on_sale,omitempty"`
}

func TestVerifyJSON(t *testing.T) {
	var p = &Product{}
	p.Name = "Phone13"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 8499.00
	p.ProductID = 0
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	fmt.Println(string(data))
}

// 结果 {"name":"Phone13","number":10000,"price":8499,"is_on_sale":true}
```

4. type，有些时候，我们在序列化或者反序列化的时候，可能结构体类型和需要的类型不一致，这个时候可以指定,支持string,number和boolean

```go
package e4

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Product 商品信息
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

func TestVerifyJSON(t *testing.T) {
	var data = `{"name":"Phone13","product_id":"1001","number":"10000","price":"8799","is_on_sale":"true"}`
	var p = new(Product)

	err := json.Unmarshal([]byte(data), &p)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	fmt.Println(*p)
}

// 结果{Phone13 1001 10000 8799 true}
```

5. Golang如何自定义解析JSON
很多时候，我们可能遇到这样的场景，就是远端返回的JSON数据不是你想要的类型，或者你想做额外的操作，比如在解析的过程中进行校验，或者类型转换，那么我们可以这样或者在解析过程中进行数据转换

```go
package custom_json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

type EMail struct {
	Value string
}

func (m *EMail) UnmarshalJSON(data []byte) error {
	if bytes.Contains(data, []byte("@")) {
		return fmt.Errorf("mail format error")
	}
	m.Value = string(data)
	return nil
}

func (m *EMail) MarshalJSON() (data []byte, err error) {
	if err != nil {
		data = []byte(m.Value)
	}
	return
}

type Phone struct {
	Value string
}

func (m *Phone) UnmarshalJSON(data []byte) error {
	if len(data) != 11 {
		return fmt.Errorf("phone format error")
	}
	m.Value = string(data)
	return nil
}

func (m *Phone) MarshalJSON() (data []byte, err error) {
	if err != nil {
		data = []byte(m.Value)
	}
	return
}

type User struct {
	Name  string
	email EMail
	phone Phone
}

func TestVerifyJSON(t *testing.T) {
	user := User{}
	user.Name = "lgq"
	user.email.Value = "lgq@126.com"
	user.phone.Value = "13588827425"
	fmt.Println(json.Marshal(user))
}
```

为什么要这样？

如果是客户端开发，需要开发大量的API，接收大量的JSON，在开发早期定义各种类型看起来是很大的工作量，不如写 if else 判断数据简单暴力。
但是到开发末期，你会发现预先定义的方式能极大的提高你的代码质量，减少代码量。
