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
