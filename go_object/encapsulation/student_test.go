package encapsulation

import (
	"fmt"
	"math/rand"
	"testing"
)

type Student struct {
	id   int
	name string
	age  int
}

func NewStudentV1(name string, age int) *Student {
	t := new(Student)
	t.id = generate()
	t.name = name
	t.age = age
	return t
}

func NewStudentV2(name string, age int) *Student {
	return &Student{id: generate(), name: name, age: age}
}

func generate() int {
	return rand.Int() + 2021
}

func (p *Student) SetName(name string) {
	p.name = name
}
func (p *Student) SetAge(age int) {
	p.age = age
}
func (p *Student) GetName() string {
	return p.name
}
func (p *Student) GetAge() int {
	return p.age
}

func TestVerify(t *testing.T)  {
	t1 := NewStudentV1("liugq",18)
	fmt.Println(t1)
	t2 := NewStudentV2("liugq",22)
	fmt.Println(t2)
}
