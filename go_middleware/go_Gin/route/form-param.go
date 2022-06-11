package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E8%B7%AF%E7%94%B1/%E8%A1%A8%E5%8D%95%E5%8F%82%E6%95%B0.html
*/
func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	r.Run()
}
