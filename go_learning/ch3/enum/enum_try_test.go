package enum

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
const (
	Open = 1 << iota
	Close
	Pending
)

func TestVerify(t *testing.T) {
	t.Log(Wednesday)
	t.Log(Pending)
}
