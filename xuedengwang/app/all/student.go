package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	globalkey "xuedengwang/core/global"
	student "xuedengwang/service/student/rpc"
)

func Student() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  globalkey.CtxConfKey,
			Value: "./deploy/etc/local/student.yaml",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Action = student.Run

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
