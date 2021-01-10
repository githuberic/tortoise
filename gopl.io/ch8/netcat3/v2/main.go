package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan string)
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println(">>>Done")
		done <- ">>>Done<<<" // 指示主goroutine
	}()
	mustCopy(conn, os.Stdin)
	defer conn.Close()
	fmt.Println(<-done) // 等待后台goroutine完成
}
