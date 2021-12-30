package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 关闭连接
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		var recvStr = string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)

		// 发送数据
		conn.Write([]byte(recvStr))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		// 启动一个goroutine处理连接
		go process(conn)
	}
}
