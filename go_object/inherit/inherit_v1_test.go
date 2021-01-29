package inherit

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) speak() {
	fmt.Printf("I'm %s,age %v", p.Name, p.Age)
}

type Programmer struct {
	Person
	Language string
}

func TestVerifyV1(t *testing.T)  {
	programmer := Programmer{
		Person :Person{
			Name: "lgq",
			Age: 35,
		},
		Language: "golang",
	}
	programmer.speak()
}