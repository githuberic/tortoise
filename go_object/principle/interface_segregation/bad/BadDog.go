package bad

import (
	"errors"
	"fmt"
)

type BadDog struct {
	iID   int
	sName string
}

func NewBadDog(id int, name string) IBadAnimal {
	return &BadDog{
		iID:   id,
		sName: name,
	}
}

func (me *BadDog) ID() int {
	return me.iID
}
func (me *BadDog) Name() string {
	return me.sName
}
func (me *BadDog) Eat() error {
	fmt.Printf("%v/%v is eating\n", me.Name(), me.ID())
	return nil
}
func (me *BadDog) Fly() error {
	return errors.New(fmt.Sprintf("%v/%v cannot fly", me.Name(), me.ID()))
}
func (me *BadDog) Swim() error {
	fmt.Printf("%v/%v is swimming\n", me.Name(), me.ID())
	return nil
}
