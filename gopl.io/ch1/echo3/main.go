package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	fmt.Printf(strings.Join(os.Args[1:]," "))
}
