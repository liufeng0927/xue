package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	globalkey "xuedengwang/core/global"
	score "xuedengwang/service/score/rpc"
)

func Scores() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  globalkey.CtxConfKey,
			Value: "./deploy/etc/local/scores.yaml",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Action = score.Run

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
