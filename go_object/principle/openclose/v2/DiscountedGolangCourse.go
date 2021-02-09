package v2

/**
该课程同时实现ICourse和IDiscount接口
*/
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
func (p *DiscountedGolangCourse) Discount() float64 {
	return p.fDiscount
}
func (p *DiscountedGolangCourse) Price() float64 {
	return p.fDiscount * p.GolangCourse.Price()
}
