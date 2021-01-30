package builder

import (
	"fmt"
	"testing"
)

func TestVerifyV1(t *testing.T) {
	//builder := new(Builder)

	director := new(Director)
	/*

	director.SetBuilder(*builder)
	*/
	computer := director.Create("Intel","Kingsidun","WD")
	fmt.Println(*computer)
}
