package reader

import (
	"fmt"
	"strings"

	"github.com/salex-org/ikea-dirigera-client/pkg/client"
)

type Device = client.Device

var readingTypeMapping = map[string]ReadingType{
	"batteryPercentage":  ReadingTypeBatteryLevel,
	"currentTemperature": ReadingTypeTemperature,
	"currentRH":          ReadingTypeHumidity,
	"currentCO2":         ReadingTypeCO2,
	"currentPM25":        ReadingTypePM25,
}

func All(client client.Client) ([]Reading, error) {
	devices, err := client.ListDevices()
	if err != nil {
		return nil, fmt.Errorf("error getting devices: %w", err)
	}

	sensors := filterReachableSensors(devices)
	roomNamesByGroupID := buildRoomNameCache(sensors)

	var readings []Reading
	processed := make(map[HashKey]struct{})

	for _, sensor := range sensors {
		attrs := sensor.Attributes
		for attrName, readingType := range readingTypeMapping {
			if value, ok := attrs[attrName]; ok {
				reading := Reading{
					Type:     readingType,
					deviceID: baseID(*sensor),
					Value:    value.(float64),
					Room:     roomNamesByGroupID[baseID(*sensor)],
				}

				if _, seen := processed[reading.HashKey()]; !seen {
					processed[reading.HashKey()] = struct{}{}
					readings = append(readings, reading)
				}
			}
		}
	}

	return readings, nil
}

// strip the suffix (_1, _2, etc..) from the deviceID
// baseID is used to match related sensors and find the room name
func baseID(sensor Device) string {
	return sensor.ID[0 : len(sensor.ID)-2]
}

func buildRoomNameCache(sensors []*Device) map[string]string {
	cache := make(map[string]string)
	for _, sensor := range sensors {
		if len(sensor.Room.Name) != 0 {
			cache[baseID(*sensor)] = strings.TrimSpace(sensor.Room.Name)
		}
	}

	return cache
}

func filterReachableSensors(devices []*Device) []*Device {
	var sensors []*Device
	for _, device := range devices {
		if device.Type != "sensor" || !device.IsReachable {
			continue
		}

		sensors = append(sensors, device)
	}

	return sensors
}
