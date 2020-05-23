package secrets

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:   "secrets",
	Usage:  "",
	Action: run,
	Flags:  []cli.Flag{},
}

func run(c *cli.Context) error {
	logrus.Info("secrets command")
	return nil
}
