package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var Exporter = &cli.Command{
	Name:        "exporter",
	Description: "start prometheus exporter server",
	Action: func(ctx context.Context, c *cli.Command) error {
		fmt.Println("nope.")
		return nil
	},
}
