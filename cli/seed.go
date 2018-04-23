package cli

import (
	"context"
	"time"

	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func InitializeSeed(c *cli.Context) (err error) {
	cli, err := client.NewEnvClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()
	if err != nil {
		return
	}
	imageName := c.Args().Get(0)
	log.Debugf("Attempting to grab '%s'", imageName)
	inspectResults, _, err := cli.ImageInspectWithRaw(ctx, imageName)
	if err != nil {
		return
	}
	for _, layer := range inspectResults.RootFS.Layers {
		log.Infof("%s", layer)
	}
	return
}
