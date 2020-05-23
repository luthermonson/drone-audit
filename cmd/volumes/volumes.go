package volumes

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:   "volumes",
	Usage:  "",
	Action: run,
	Flags:  []cli.Flag{},
}

func run(c *cli.Context) error {
	logrus.Info("volumes command")
	return nil
}
