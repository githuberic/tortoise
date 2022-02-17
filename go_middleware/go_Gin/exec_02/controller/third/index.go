package third

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.POST("/third", geFunction)
	return
}

func geFunction(c *gin.Context) {
	fmt.Println("third .........")
	return
}
