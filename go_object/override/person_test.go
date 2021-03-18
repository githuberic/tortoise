package override

import "testing"

func TestVerify(t *testing.T)  {
	m := &Man{}
	m.Eat()
	m.Run()
	m.Sleep()
}
