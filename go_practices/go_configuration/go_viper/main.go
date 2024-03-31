package main

import (
	"errors"
	"github.com/spf13/pflag"
	"go_practices/go_configuration/go_viper/config"
	"go_practices/go_configuration/go_viper/model"
	"log"
)

var (
	conf = pflag.StringP("config", "c", "", "config filepath")
)

func main() {
	pflag.Parse()

	// 初始化配置
	if err := config.Init(*conf); err != nil {
		panic(err)
	}

	// 连接mysql数据库
	isSuc := model.GetInstance().InitPool()
	if !isSuc {
		log.Println("init database pool failure...")
		panic(errors.New("init database pool failure"))
	}
}
