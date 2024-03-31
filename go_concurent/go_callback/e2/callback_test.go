package e2

import (
	"fmt"
	"strconv"
	"testing"
)

type Callback func(msg string)

func stringToInt(s string, callback Callback) int64 {
	if v, err := strconv.ParseInt(s, 0, 0); err != nil {
		callback(err.Error())
		return 0
	} else {
		return v
	}
}

func errLog(msg string) {
	fmt.Println("Convert error: ", msg)
}

func TestCallback(t *testing.T) {
	fmt.Println(stringToInt("18", errLog))
	fmt.Println(stringToInt("hh", errLog))
}
