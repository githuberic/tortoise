package main

/*
 main函数将一个处理函数和以/开头的url链接在一起；所有的url都使用这个函数处理，启动服务器监听进入8000端口的请求
 一个请求由一个http.request类型的结构体表示

运行方式：浏览器运行 http://localhost:8000
*/
import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}
