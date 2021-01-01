package panic_recover_test

import (
	"errors"
	"fmt"
	"testing"
)

// let it crash,往往是应对不确定错误的最好方法
func TestPanicVxExit(t *testing.T) {
	defer func() {
		// 错误恢复，程序没有退出
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
	//fmt.Println("End")
}

