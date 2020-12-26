package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	// 两个全局通道entering/leaving 通知客户的到来和离开
	entering = make(chan client)
	leaving  = make(chan client)
	// all incoming client messages
	messages = make(chan string)
)

// broadcaster goroutine
// 广播器broadcaster是关于如何使用select的一个规范说明,它需要对三种不同的消息进行响应
func broadcaster() {
	// 使用clients记录当前conn 集合,每个客户唯一被记录的信息是其对外发送消息的通道ID
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			// 删除&关闭连接
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

//!+handleConn
// 每个连接都有连接处理(handleConn)
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

// 客户端写入go routine
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main 主goroutine
func main() {
	// 主go routine是监听端口,接收客户端的的网络连接，
	// 对于每一个conn都创建一个新的handleConn go routine
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
