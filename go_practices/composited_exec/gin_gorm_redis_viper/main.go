package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_practices/composited_exec/gin_gorm_redis_viper/global"
	"go_practices/composited_exec/gin_gorm_redis_viper/router"
	"log"
)

//init
func init() {
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = global.SetupDBLink()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = global.SetupRedisDb()
	if err != nil {
		log.Fatalf("init.SetupRedisDb err: %v", err)
	}
}

func main() {
	//设置运行模式
	gin.SetMode(global.ServerSetting.RunMode)
	//引入路由
	r := router.Router()
	//run
	r.Run(":" + global.ServerSetting.HttpPort)
}
