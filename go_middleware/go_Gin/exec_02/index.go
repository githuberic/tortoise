package main

import (
	"github.com/gin-gonic/gin"
	c "tortoise/go_middleware/go_Gin/exec_02/controller"
)

func main() {
	r := gin.Default()
	rr := c.GinRouter(r)

	// 监听并在 0.0.0.0:8080 上启动服务
	rr.Run(":8080")
}
