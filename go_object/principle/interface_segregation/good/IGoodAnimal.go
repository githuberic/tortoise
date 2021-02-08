package good

type IGoodAnimal interface {
	ID() int
	Name() string
}

type ISupportEat interface {
	Eat() error
}

type ISupportFly interface {
	Fly() error
}

type ISupportSwim interface {
	Swim() error
}
