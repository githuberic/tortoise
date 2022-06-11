package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tortoise/go_middleware/go_Gin/exec_02/controller/forth"
	"tortoise/go_middleware/go_Gin/exec_02/controller/second"
	"tortoise/go_middleware/go_Gin/exec_02/controller/third"
)

func GinRouter(r *gin.Engine) *gin.Engine {
	rr := r.Group("/")
	rr.GET("/first", func(c *gin.Context) {
		fmt.Println("first .........")
	})

	rr = r.Group("/a")
	second.Routers(rr)
	third.Routers(rr)
	forth.Routers(rr)

	return r
}
