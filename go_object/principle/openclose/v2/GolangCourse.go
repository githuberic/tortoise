package v2

/**
golang课程类, 实现ICourse接口
*/
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
func (p *GolangCourse) Name() string {
	return p.sName
}
func (p *GolangCourse) Price() float64 {
	return p.fPrice
}
