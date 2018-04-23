package main

import (
	"os"

	planter "github.com/seemethere/planter/cli"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := planter.InitializeCLI()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
