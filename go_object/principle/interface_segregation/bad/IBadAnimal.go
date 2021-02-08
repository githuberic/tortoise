package bad

type IBadAnimal interface {
	ID() int
	Name() string

	Eat() error
	Fly() error
	Swim() error
}
