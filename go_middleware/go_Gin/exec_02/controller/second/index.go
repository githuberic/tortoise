package second

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.POST("/second", Function)
	return
}

func Function(c *gin.Context) {
	fmt.Println("second .........")
	return
}
