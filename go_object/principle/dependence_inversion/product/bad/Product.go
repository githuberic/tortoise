package bad

/**
拍品
*/
type Product struct {
	id    int
	name  string
	price float32
}

func NewProduct(id int, name string, price float32) *Product {
	return &Product{
		id:    id,
		name:  name,
		price: price,
	}
}

func (p *Product) CalcAuthorPrice() {
	p.price = p.price * 0.88
}

func (p *Product) CalcBrandPrice() {
	p.price = p.price * 0.80
}

func (p *Product) CalcAuctionPrice() {
	p.price = p.price * 0.67
}
