package images

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:   "images",
	Usage:  "",
	Action: run,
	Flags:  []cli.Flag{},
}

func run(c *cli.Context) error {
	logrus.Info("images command")
	return nil
}
