package reader

type ReadingType string

const (
	ReadingTypeTemperature  = "temperature"
	ReadingTypeHumidity     = "humidity"
	ReadingTypeCO2          = "co2"
	ReadingTypePM25         = "pm25"
	ReadingTypeBatteryLevel = "battery-level"
)

type Reading struct {
	Type  ReadingType
	Room  string
	Value float64

	deviceID string
}

type HashKey struct {
	Type    ReadingType
	GroupID string
}

func (r Reading) HashKey() HashKey {
	return HashKey{
		Type:    r.Type,
		GroupID: r.deviceID,
	}
}
