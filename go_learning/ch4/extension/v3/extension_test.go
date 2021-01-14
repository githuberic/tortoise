package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}
func (p *Pet) Speak() {
	fmt.Print("...")
}
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}
/*
func (d *Dog) Speak() {
	d.p.Speak()
	fmt.Print("Wang!")
}
func (d *Dog) SpeakTo(host string) {
	d.p.SpeakTo(host)
}*/
func TestDogV2(t *testing.T) {
	dog := new(Dog)
	dog.p.Speak()
	dog.p.SpeakTo("Chao")
}
