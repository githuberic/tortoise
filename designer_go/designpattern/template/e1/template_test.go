package e1

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	football := NewFootBall()
	football.Play()

	fmt.Println("-------------------")
	basketball := NewBasketball()
	basketball.Play()
}
