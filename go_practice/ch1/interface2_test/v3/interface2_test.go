package interface2_test

import (
	"fmt"
	"testing"
)

type Country struct {
	Name string
}

type City struct {
	Name string
}

/**
使用了一个叫Stringable 的接口，我们用这个接口把“业务类型” Country 和 City 和“控制逻辑” Print() 给解耦了。
*/
type Stringable interface {
	ToString() string
}

func (c Country) ToString() string {
	return "Country = " + c.Name
}
func (c City) ToString() string {
	return "City = " + c.Name
}

/**
只要实现了Stringable 接口，都可以传给 PrintStr() 来使用。
*/
func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}

func TestVerify(t *testing.T) {
	d1 := Country{"China"}
	d2 := City{"HangZhou"}
	PrintStr(d1)
	PrintStr(d2)
}
