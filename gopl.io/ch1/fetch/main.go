package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/**
获取一个URL
 */
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			// 进程退出时，返回状态码1
			os.Exit(1)
		}

		// ioutil.ReadAll 读取整个response并存入b,关闭body数据流来避免资源泄露
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
