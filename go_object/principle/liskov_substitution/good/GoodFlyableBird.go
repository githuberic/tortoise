package good

import "fmt"

type GoodFlyableBird struct {
	GoodNormalBird
}

func NewGoodFlyableBird(id int, name string) IGoodBird {
	return &GoodFlyableBird{
		*NewGoodNormalBird(id, name),
	}
}
func (p *GoodFlyableBird) Fly() error {
	fmt.Printf("GoodFlyableBird.Fly, id=%v, name=%v\n", p.ID(), p.Name())
	return nil
}
