package good

import "fmt"

type GoodBird struct {
	GoodAnimalInfo
}

func NewGoodBird(id int,name string) IGoodAnimal {
	return &GoodBird{
		GoodAnimalInfo{
			id,
			name,
		},
	}
}

func (me *GoodBird) Eat() error {
	fmt.Printf("%v/%v is eating\n", me.Name(), me.ID())
	return nil
}

func (me *GoodBird) Fly() error {
	fmt.Printf("%v/%v is flying\n", me.Name(), me.ID())
	return nil
}
