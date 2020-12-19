/*
 go routine 是一个并发执行的函数
 通道：是一种允许某一例程向另外一个例程传递指定类型的值的通信机制
 main 函数在一个goroutine中执行，go 语句创建额外的goroutine
 */
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// 发送到通道ch
		ch <- fmt.Sprint(err)
		return
	}

	// io.copy函数读取响应内容，写入 ioutil.Discard输出流进行丢弃,copy返回字节数以及任何出现的错误
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	// 防止资源泄露
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	// 每一个结果出现时，fetch发送一行汇总信息到通道ch，main函数中的第二轮循环接收并且输出汇总行
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		// 启动一个goroutine
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // 从通道ch接收
	}

	fmt.Printf("%.2f elapsed \n", time.Since(start).Seconds())
}
