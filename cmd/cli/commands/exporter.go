package commands

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tmw/pi-climate-monitor/internal/client"
	"github.com/tmw/pi-climate-monitor/internal/reader"
	"github.com/urfave/cli/v3"
)

var Exporter = &cli.Command{
	Name:        "exporter",
	Description: "start prometheus exporter server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "access-token",
			Usage:    "access token to be used when fetching the readout",
			Required: true,
			Sources:  cli.EnvVars("ACCESS_TOKEN"),
		},
		&cli.StringFlag{
			Name:     "port",
			Usage:    "port the prometheus exporter is listening on",
			Required: false,
			Value:    "9000",
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

		e := exporter{client: client}
		r := prometheus.NewRegistry()
		r.MustRegister(e)
		http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
		return http.ListenAndServe(":"+c.String("port"), nil)
	},
}

var metric = prometheus.NewDesc(
	"climate",
	"reading from a climate sensor",
	[]string{"sensor_type", "room"},
	nil,
)

type exporter struct {
	client client.Client
}

func (e exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- metric
}

func (e exporter) Collect(ch chan<- prometheus.Metric) {
	readings, err := reader.All(e.client)
	if err != nil {
		panic(err)
	}

	for _, reading := range readings {
		ch <- prometheus.MustNewConstMetric(
			metric,
			prometheus.GaugeValue,
			reading.Value,
			string(reading.Type), reading.Room,
		)
	}

}
