package v2

import "testing"

type ICourse interface {
	ID() int
	Name() string
	Price() float64
}

type GolangCourse struct {
	iID    int
	sName  string
	fPrice float64
}

func NewGolangCourse(id int, name string, price float64) ICourse {
	return &GolangCourse{
		iID:    id,
		sName:  name,
		fPrice: price,
	}
}
func (p *GolangCourse) ID() int {
	return p.iID
}
func (this *GolangCourse) Name() string {
	return this.sName
}
func (this *GolangCourse) Price() float64 {
	return this.fPrice
}

type IDiscount interface {
	Discount() float64
}

type DiscountedGolangCourse struct {
	GolangCourse
	fDiscount float64
}

func NewDiscountGolangCourse(id int, name string, price float64, discount float64) ICourse {
	return &DiscountedGolangCourse{
		GolangCourse: GolangCourse{
			iID:    id,
			sName:  name,
			fPrice: price,
		},
		fDiscount: discount,
	}
}

func (this *DiscountedGolangCourse) Discount() float64 {
	return this.fDiscount
}

func (this *DiscountedGolangCourse) Price() float64  {
	return this.fDiscount * this.GolangCourse.Price()
}

func TestVerify(t *testing.T)  {
	fnShowCounse := func(it ICourse) {
		t.Logf("id=%v, name=%v, price=%v\n", it.ID(), it.Name(), it.Price())
	}

	c1 := NewGolangCourse(1, "golang课程", 100)
	fnShowCounse(c1)

	c2 := NewDiscountGolangCourse(2, "golang优惠课程", 100, 0.6)
	fnShowCounse(c2)
}

// https://studygolang.com/articles/33100