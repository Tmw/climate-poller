package commands

import (
	"context"
	"fmt"

	"github.com/tmw/pi-climate-monitor/internal/client"
	"github.com/urfave/cli/v3"
)

var Authorize = &cli.Command{
	Name: "authorize",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "client-name",
			Value: "climate-poller",
		},
	},
	Description: "Retrieve access token from Dirigera hub",
	Action: func(ctx context.Context, c *cli.Command) error {
		hub, err := client.FirstHub()
		if err != nil {
			return fmt.Errorf("error getting hub: %w", err)
		}

		auth, err := client.GetAccessToken(*hub, c.String("client-name"))
		if err != nil {
			return fmt.Errorf("error getting access token: %w", err)
		}

		encoded, err := client.StringFromAuth(*auth)
		if err != nil {
			return fmt.Errorf("error encoding access token: %w", err)
		}

		fmt.Print(encoded)
		return nil
	},
}
