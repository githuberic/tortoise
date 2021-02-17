package encapsulation

type Person struct {
	name string
	age  int
}

func (p *Person) SetName(name string) {
	p.name = name
}
func (p *Person) SetAge(age int) {
	p.age = age
}
func (p *Person) GetName() string {
	return p.name
}
func (p *Person) GetAge() int {
	return p.age
}
