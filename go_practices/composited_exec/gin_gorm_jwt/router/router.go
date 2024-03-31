package router

import (
	"github.com/gin-gonic/gin"
	"go_practices/composited_exec/gin_gorm_jwt/controller"
	"go_practices/composited_exec/gin_gorm_jwt/global"
	"go_practices/composited_exec/gin_gorm_jwt/middleware"
	"go_practices/composited_exec/gin_gorm_jwt/pkg/result"
	"runtime/debug"
)

func Router() *gin.Engine {
	router := gin.Default()
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(middleware.AccessLog())
	router.Use(Recover)

	// 路径映射
	articlec := controller.NewArticleController()
	router.GET("/article/getone/:id", articlec.GetOne)
	router.GET("/article/list", articlec.GetList)

	userc := controller.NewUserController()
	router.GET("/user/login", userc.Login)
	router.GET("/user/info", middleware.JWTAuthMiddleware(), userc.Info)
	router.GET("/user/pass", userc.Pass)
	return router
}

//404
func HandleNotFound(c *gin.Context) {
	global.Logger.Errorf("handle not found: %v", c.Request.RequestURI)
	result.NewResult(c).Error(404, "资源未找到")
	return
}

//500
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			//log.Printf("panic: %v\n", r)
			global.Logger.Errorf("panic: %v", r)
			//log stack
			global.Logger.Errorf("stack: %v", string(debug.Stack()))
			debug.PrintStack()
			result.NewResult(c).Error(500, "服务器内部错误")
		}
	}()
	//继续后续接口调用
	c.Next()
}
