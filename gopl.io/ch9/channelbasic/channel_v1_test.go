package channelbasic

import (
	"fmt"
	"testing"
)

func TestVerifyV1(t *testing.T) {
	ch := incomingURLs()
	for str := range ch{
		fmt.Println(str)
	}

	/*
	for str := range incomingURLs(){
		fmt.Println(str)
	}*/
}

func incomingURLs() <-chan string {
	ch := make(chan string)

	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io"} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

func incomingURLsv1() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}
