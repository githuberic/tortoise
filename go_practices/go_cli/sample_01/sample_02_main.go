package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

/**
run example
1:go run main.go --port 7777
2:go run ./sample_02_main.go --port 7777 --host 192.168.0.1 --debug true
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

	//预置变量
	var host string
	var debug bool

	// 设置启动参数
	oApp.Flags = []cli.Flag{
		//参数类型string,int,bool
		cli.StringFlag{
			Name:        "host",           //参数名字
			Value:       "127.0.0.1",      //参数默认值
			Usage:       "Server Address", //参数功能描述
			Destination: &host,            //接收值的变量
		},
		cli.IntFlag{
			Name:  "port,p",
			Value: 8888,
			Usage: "Server port",
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "debug mode",
			Destination: &debug,
		},
	}

	//该程序执行的代码
	oApp.Action = func(c *cli.Context) error {
		fmt.Printf("host=%v \n", host)
		fmt.Printf("port=%v \n", c.Int("port")) //不使用变量接收，直接解析
		fmt.Printf("debug=%v \n", debug)
		return nil
	}

	//启动
	if err := oApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
