package v2

type ICourse interface {
	ID() int
	Name() string
	Price() float64
}
