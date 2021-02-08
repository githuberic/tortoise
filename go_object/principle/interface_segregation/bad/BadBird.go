package bad

import (
	"errors"
	"fmt"
)

type BadBird struct {
	iID   int
	sName string
}

func NewBadBird(id int, name string) IBadAnimal {
	return &BadBird{
		iID:   id,
		sName: name,
	}
}

func (p *BadBird) ID() int {
	return p.iID
}
func (p *BadBird) Name() string {
	return p.sName
}
func (p *BadBird) Eat() error {
	fmt.Printf("%v/%v is eating\n", p.Name(), p.ID())
	return nil
}
func (p *BadBird) Fly() error {
	fmt.Printf("%v/%v is flying\n", p.Name(), p.ID())
	return nil
}
func (me *BadBird) Swim() error {
	return errors.New(fmt.Sprintf("%v/%v cannot swimming", me.Name(), me.ID()))
}
