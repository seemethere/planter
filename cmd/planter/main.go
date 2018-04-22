package main

import (
	"os"

	planterCli "github.com/seemethere/planter/cli"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Planter"

	app.Action = planterCli.InitializeSeed
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
