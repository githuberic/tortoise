package v2

/**
课程接口
*/
type ICourse interface {
	ID() int
	Name() string
	Price() float64
}
