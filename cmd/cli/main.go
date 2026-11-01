package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tmw/pi-climate-monitor/cmd/cli/commands"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := cli.Command{
		Name:  "climate-poller",
		Usage: "CLI tool to poll Dirigera connected smart devices and manage access tokens",
		Commands: []*cli.Command{
			{
				Name: "create-access-token",
				Action: func(ctx context.Context, c *cli.Command) error {
					fmt.Println("unimplemented")
					return nil
				},
			},

			commands.Dump,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
