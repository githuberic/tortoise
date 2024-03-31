package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ms "go_practices/go_gin/exec_03/model"
)

func Function(c *gin.Context, input ms.Input) {
	fmt.Println("second .........,input:", input.Id)
	return
}
