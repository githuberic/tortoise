package v1

// 结构体可以使用嵌套匿名结构体所有的字段和方法
// 首字母大写或者小写的字段、方法，都可以使用。
import (
	"fmt"
	"testing"
)

/**
拍品
*/
type Product struct {
	Name  string
	price float32
}
func (p *Product) String() {
	fmt.Printf("Product %s,price %v", p.Name, p.price)
}

/**
品牌馆
*/
type BrandProduct struct {
	Product
	brand string
}

func TestVerifyV1(t *testing.T) {
	programmer := BrandProduct{
		Product: Product{
			Name: "翔龙武士",
			price:  988,
		},
		brand: "龙泉宝剑",
	}
	// 匿名结构体字段访问可以简化
	programmer.String()

	// 首字母大写或者小写的字段、方法，都可以使用。
	//programmer.Product.String()
}
