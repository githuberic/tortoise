package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

/**
to do ... outer
 */
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5}
	reverse(arr[:])
	fmt.Println(arr)

	arr1 := []int{0, 1, 2, 3, 4, 5}
	reverse(arr1[:2])
	fmt.Println(arr1)

	reverse(arr1[2:])
	fmt.Println(arr1)

	reverse(arr1)
	fmt.Println(arr1) // "[2 3 4 5 0 1]"

	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
}
