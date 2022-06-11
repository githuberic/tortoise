package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
curl -l 'http://localhost:8000/user?name=lgq'
*/
func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8080/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "枯藤")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	//默认为监听8080端口
	r.Run(":8000")
}
