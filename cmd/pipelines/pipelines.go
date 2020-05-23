package pipelines

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:   "pipelines",
	Usage:  "",
	Action: run,
	Flags:  []cli.Flag{},
}

func run(c *cli.Context) error {
	logrus.Info("pipelines command")
	return nil
}
