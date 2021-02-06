package bad

import "fmt"

// BadNormalBird实现了IBadBird接口
type BadNormalBird struct {
	iID   int
	sName string
}

func NewBadNormalBird(id int, name string) IBadBird {
	return &BadNormalBird{
		id,
		name,
	}
}
func (me *BadNormalBird) ID() int {
	return me.iID
}
func (me *BadNormalBird) Name() string {
	return me.sName
}
func (me *BadNormalBird) Tweet() error {
	fmt.Printf("BadNormalBird.Tweet, id=%v, name=%v\n", me.ID(), me.Name())
	return nil
}
func (me *BadNormalBird) Fly() error {
	fmt.Printf("BadNormalBird.Fly, id=%v, name=%v\n", me.ID(), me.Name())
	return nil
}
