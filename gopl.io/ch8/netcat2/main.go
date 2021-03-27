package main

import (
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

	defer conn.Close()

	// 从conn到console输出
	go mustCopy(os.Stdout, conn)

	// 从console输入copy到conn
	mustCopy(conn, os.Stdin)
}
