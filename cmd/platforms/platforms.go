package platforms

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:   "platforms",
	Usage:  "",
	Action: run,
	Flags:  []cli.Flag{},
}

func run(c *cli.Context) error {
	logrus.Info("platforms command")
	return nil
}
