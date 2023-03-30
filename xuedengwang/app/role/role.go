package main

import (
	"github.com/urfave/cli/v2"
	"log"
	globalkey "xuedengwang/core/global"
	role "xuedengwang/service/role/rpc"

	"os"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  globalkey.CtxConfKey,
			Value: "./deploy/etc/local/role.yaml",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Action = role.Run

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
