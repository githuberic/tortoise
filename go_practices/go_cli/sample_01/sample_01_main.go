package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

/**
go run main.go help
*/
func main() {
	//实例化一个命令行程序
	oApp := cli.NewApp()
	//程序名称
	oApp.Name = "GoTool"
	//程序的用途描述
	oApp.Usage = "To save the world"
	//程序的版本号
	oApp.Version = "1.0.0"
	//该程序执行的代码
	oApp.Action = func(c *cli.Context) error {
		fmt.Println("sample_01_main")
		return nil
	}
	//启动
	if err := oApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
