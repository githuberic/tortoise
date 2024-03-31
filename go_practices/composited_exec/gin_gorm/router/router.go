package router

import (
	"github.com/gin-gonic/gin"
	"go_practices/composited_exec/gin_gorm/controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/article/getone/:id", controller.NewArticleController().GetOne)
	router.GET("/article/list", controller.NewArticleController().GetList)
	return router
}
