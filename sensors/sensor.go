package sensors

import (
	"encoding/json"
	"time"
)

type Sensor struct {
	Id         string
	Value      int
	RecordedAt time.Time
}

func (sensor *Sensor) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Id         string `json:"id"`
			Value      int    `json:"value"`
			RecordedAt int64  `json:"read_at"`
		}{
			Id:         sensor.Id,
			Value:      sensor.Value,
			RecordedAt: sensor.RecordedAt.Unix(),
		},
	)
}
