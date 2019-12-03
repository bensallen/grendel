package cmd

import (
	"github.com/urfave/cli"
)

func NewHostCommand() cli.Command {
	return cli.Command{
		Name:        "host",
		Usage:       "Host commands",
		Description: "Host commands",
		Subcommands: []cli.Command{
			NewNetbootCommand(),
		},
	}
}
