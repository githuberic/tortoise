package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	ages["eric"] = 35

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	// 判定是否存在
	age, ok := ages["eric"]
	if ok {
		fmt.Printf("bob(age=%d) exist\n", age)
	} else {
		fmt.Println("not exist")
	}
}
