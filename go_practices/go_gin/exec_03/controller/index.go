package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ms "go_practices/go_gin/exec_03/model"
	ss "go_practices/go_gin/exec_03/services"
)

func Routers(r *gin.RouterGroup) {
	rr := r.Group("")
	rr.POST("/second", Function)
	return
}
func Function(c *gin.Context) {
	var input ms.Input
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
	}
	ss.Function(c, input)
	return
}
