package cli

import (
	"context"
	"time"

	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func InitializeSeed(c *cli.Context) (err error) {
	client, err := client.NewEnvClient()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()
	if err != nil {
		return
	}
	imageName := c.Args().Get(0)
	log.Debugf("Attempting to grab '%s'", imageName)
	layers, err := GrabImageLayers(ctx, client, imageName)
	if err != nil {
		return
	}
	for _, layer := range layers {
		log.Info(layer)
	}
	return
}

// Grab the layers for a given image
func GrabImageLayers(ctx context.Context, client *client.Client, imageName string) (layers []string, err error) {
	inspectResults, _, err := client.ImageInspectWithRaw(ctx, imageName)
	if err != nil {
		return
	}
	layers = inspectResults.RootFS.Layers
	return
}
