package good

import "fmt"

type GoodDog struct {
	GoodAnimalInfo
}

func NewGoodDog(id int, name string) IGoodAnimal {
	return &GoodDog{
		GoodAnimalInfo{
			id,
			name,
		},
	}
}

func (me *GoodDog) Eat() error {
	fmt.Printf("%v/%v is eating\n", me.Name(), me.ID())
	return nil
}

func (me *GoodDog) Swim() error {
	fmt.Printf("%v/%v is swimming\n", me.Name(), me.ID())
	return nil
}