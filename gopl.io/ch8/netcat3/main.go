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

	done := make(chan struct{})
	go func() {
		fmt.Println(">>>Start")
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Println("io.copy error")
		}
		fmt.Println(">>>Done")
		done <- struct{}{} // 指示主goroutine
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // 等待后台goroutine完成
}
