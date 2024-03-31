package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_practices/composited_exec/gin_gorm_tx/global"
	"go_practices/composited_exec/gin_gorm_tx/router"
	"log"
)

func init() {
	err := global.SetupDBLink()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	//引入路由
	r := router.Router()
	//run
	r.Run(":8080")
}
