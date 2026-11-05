package main

import (
	"context"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tmw/pi-climate-monitor/cmd/cli/commands"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := cli.Command{
		Name:  "climate-poller",
		Usage: "CLI tool to poll Dirigera connected smart devices and manage access tokens",
		Commands: []*cli.Command{
			commands.Authorize,
			commands.Exporter,
			commands.Dump,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
