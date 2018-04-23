package cli

import (
	"github.com/urfave/cli"
)

func InitializeCLI() (app *cli.App) {
	app = cli.NewApp()
	app.Name = "Planter"
	app.Commands = []cli.Command{
		{
			Name:   "seed",
			Usage:  "seed an image from this machine",
			Action: InitializeSeed,
		},
	}
	return
}
