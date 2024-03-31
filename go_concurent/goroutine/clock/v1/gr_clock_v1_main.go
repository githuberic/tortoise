package main

import (
	"io"
	"log"
	"net"
	"time"
)

/**
时钟服务器,每秒一次向客户端发送当前时间
go build ./main
./main &

访问的客户端 nc localhost 8000
 */
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
	// listen 函数创建一个net.listener对象，，监听连接
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// 监听器的accept被阻塞，直到有连接请求进来
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn) // 一次处理一个连接
	}
}
