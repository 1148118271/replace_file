package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"replace_file/util"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "config, c",
			Usage: "配置文件, 为空则默认取当前文件夹下config.conf文件",
		},
	}
	app.Action = func(c *cli.Context){
		log.Println("开始替换文件, 程序开始.....")
		conf := c.String("config")
		fmt.Println(conf)
		if "" == conf {
			conf = "config.conf"
		}
		config, err := util.GetConfig(conf)
		if err != nil {
			log.Fatalln(err)
		}
		err = util.CopyFile(config)
		if err != nil {
			log.Fatalln(err)
		}
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("命令行运行异常:", err)
	}
	log.Println("所有文件替换成功, 程序结束.....")
}
