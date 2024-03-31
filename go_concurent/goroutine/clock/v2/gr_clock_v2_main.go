package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		// 服务器是顺序执行的，一个只能处理一个客户请求,在调用handleConn增加一个go, 让它在自己的go routine里面运行
		// 并发处理连接
		go handleConn(conn)
	}
}

