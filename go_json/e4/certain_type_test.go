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
