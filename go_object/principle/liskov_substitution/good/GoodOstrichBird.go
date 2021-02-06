package good

import "fmt"

type GoodOstrichBird struct {
	GoodNormalBird
}

func NewGoodOstrichBird(id int, name string) IGoodBird {
	return &GoodOstrichBird{
		*NewGoodNormalBird(id, name),
	}
}
func (me *GoodOstrichBird) Run() error {
	fmt.Printf("GoodOstrichBird.Run, id=%v, name=%v\n", me.ID(), me.Name())
	return nil
}
