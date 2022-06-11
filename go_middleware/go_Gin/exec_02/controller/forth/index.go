package forth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.POST("/forth", Function1)
	rr.POST("/five", Function2)
	return
}

func Function1(c *gin.Context) {
	fmt.Println("forth .........")
	return
}
func Function2(c *gin.Context) {
	fmt.Println("five .........")
	return
}
