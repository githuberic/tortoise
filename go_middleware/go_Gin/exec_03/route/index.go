package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tortoise/go_middleware/go_Gin/exec_03/controller"
)

func GinRouter(r *gin.Engine) *gin.Engine {
	rr := r.Group("/")
	rr.GET("/first", func(c *gin.Context) {
		fmt.Println("first .........")
	})

	rr = r.Group("/a")
	controller.Routers(rr)

	return r
}
