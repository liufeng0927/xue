package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	globalkey "xuedengwang/core/global"
	gradeclass "xuedengwang/service/gradeclass/rpc"
)

func Gradeclass() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  globalkey.CtxConfKey,
			Value: "./deploy/etc/local/gradeclass.yaml",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Action = gradeclass.Run

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
