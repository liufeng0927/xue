package main

import (
	"github.com/urfave/cli/v2"
	"log"
	globalkey "xuedengwang/core/global"
	user "xuedengwang/service/user/rpc"

	"os"
)

func User() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  globalkey.CtxConfKey,
			Value: "./deploy/etc/local/user.yaml",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Action = user.Run

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
