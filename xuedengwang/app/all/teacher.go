package main

import (
	"github.com/urfave/cli/v2"
	"log"
	globalkey "xuedengwang/core/global"
	teacher "xuedengwang/service/teacher/rpc"

	"os"
)

func Teacher() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  globalkey.CtxConfKey,
			Value: "./deploy/etc/local/teacher.yaml",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Action = teacher.Run

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
