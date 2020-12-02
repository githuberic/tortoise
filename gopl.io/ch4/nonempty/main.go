package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, v := range strings {
		if v != "" {
			strings[i] = v
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, str := range strings {
		if str != "" {
			out = append(out, str)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)
}
