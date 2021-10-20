package go_runtime

import (
	"fmt"
	"runtime"
	"testing"
)

func TestRuntimeBasic(t *testing.T)  {
	fmt.Println("cpus:",runtime.NumCPU())
	fmt.Println("goroot:",runtime.GOROOT())
	fmt.Println("archieve:",runtime.GOOS)
}
