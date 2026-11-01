package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
	"text/tabwriter"

	"github.com/tmw/pi-climate-monitor/internal/client"
	"github.com/tmw/pi-climate-monitor/internal/reader"
	"github.com/urfave/cli/v3"
)

var Dump = &cli.Command{
	Name:        "dump",
	Description: "One time readout of reachable sensors",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "format",
			Usage: "format of the dump (json, table)",
			Value: "json",
			Validator: func(s string) error {
				valid := []string{"json", "table"}
				if !slices.Contains(valid, s) {
					return fmt.Errorf(
						"invalid flag value: %s, must be one of: %s",
						s,
						strings.Join(valid, ", "),
					)
				}

				return nil
			},
		},
		&cli.StringFlag{
			Name:     "access-token",
			Usage:    "access token to be used when fetching the readout",
			Required: true,
			Sources:  cli.EnvVars("ACCESS_TOKEN"),
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		hub, err := client.FirstHub()
		if err != nil {
			return fmt.Errorf("error getting hub: %w", err)
		}

		client, err := client.NewClient(hub, c.String("access-token"))
		if err != nil {
			return fmt.Errorf("error building client: %w", err)
		}

		readings, err := reader.All(client)
		if err != nil {
			return fmt.Errorf("error getting readings: %w", err)
		}

		format := strings.ToLower(c.String("format"))
		if format == "json" {
			formatJSON(readings)
			return nil
		}

		if format == "table" {
			formatTable(readings)
			return nil
		}

		return nil
	},
}

func formatJSON(readings []reader.Reading) {
	json.NewEncoder(os.Stdout).Encode(readings)
}

func formatTable(readings []reader.Reading) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "Room\tSensor Type\tValue")
	fmt.Fprintln(w, "----\t-----------\t-----")
	for _, reading := range readings {
		fmt.Fprintf(w, "%s\t%s\t%f\n", reading.Room, reading.Type, reading.Value)
	}

	w.Flush()
}
