package main

import (
	"os"

	"github.com/seemethere/planter/commands"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Planter"

	app.Action = commands.InitializeSeed
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
