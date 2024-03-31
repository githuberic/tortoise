package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	// Flags 用于设置参数。
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}

	// Action 对应的函数就是你具体对各个参数具体的处理逻辑。
	app.Action = func(c *cli.Context) error {
		name := "Nefertiti"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}

		if c.String("lang") == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}

	app.Run(os.Args)
}
