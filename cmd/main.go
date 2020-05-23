package cmd

import (
	"github.com/luthermonson/drone-audit/cmd/images"
	"github.com/luthermonson/drone-audit/cmd/pipelines"
	"github.com/luthermonson/drone-audit/cmd/platforms"
	"github.com/luthermonson/drone-audit/cmd/secrets"
	"github.com/luthermonson/drone-audit/cmd/volumes"
	"github.com/urfave/cli/v2"
)

func Commands() cli.Commands {
	return cli.Commands{
		images.Command,
		pipelines.Command,
		platforms.Command,
		secrets.Command,
		volumes.Command,
	}
}
