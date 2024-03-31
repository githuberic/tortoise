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

	// 主goroutine从标准输入读取并发送服务器，第二个goroutine读取服务器的回复并且输出
	// 从console输入copy到conn
	mustCopy(conn, os.Stdin)
}
