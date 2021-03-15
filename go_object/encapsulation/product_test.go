package encapsulation

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
拍品
*/
type Product struct {
	id    int
	name  string
	price float32
}

/**
构造函数
*/
func NewProductV1(name string, price float32) *Product {
	t := new(Product)
	t.id = generate()
	t.name = name
	t.price = price
	return t
}

/**
构造函数
*/
func NewProductV2(name string, price float32) *Product {
	return &Product{id: generate(), name: name, price: price}
}

/**
Id生成器
*/
func generate() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

func (p *Product) SetName(name string) {
	p.name = name
}
func (p *Product) SetPrice(price float32) {
	p.price = price
}

func (p *Product) GetName() string {
	return p.name
}
func (p *Product) GetPrice() float32 {
	return p.price
}

func TestVerify(t *testing.T) {
	t1 := NewProductV1("新疆和田玉", 18000.00)
	fmt.Println(t1)

	t2 := NewProductV2("羊脂玉精品级", 18098.00)
	fmt.Println(t2)
}
