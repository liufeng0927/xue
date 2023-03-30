package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	globalkey "xuedengwang/core/global"
	gateway "xuedengwang/service/gateway/api"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  globalkey.CtxConfKey,
			Value: "./deploy/etc/local/gateway.yaml",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Action = gateway.Run

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
