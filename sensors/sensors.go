package sensors

import (
	"encoding/json"
	"time"
)

var sensors []Sensor

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

func GetSensorById(id string) *Sensor {
	for _, sensor := range sensors {
		if sensor.Id == id {
			return &sensor
		}
	}
	return nil
}

func RecordValue(id string, value int, recordedAt time.Time) {
	existingIndex, ok := findSensorIndexById(id)
	if ok {
		sensors[existingIndex].Value = value
		sensors[existingIndex].RecordedAt = recordedAt
	} else {
		newSensor := Sensor{
			Id:         id,
			Value:      value,
			RecordedAt: recordedAt,
		}
		sensors = append(sensors, newSensor)
	}
}

func findSensorIndexById(id string) (int, bool) {
	for index, sensor := range sensors {
		if sensor.Id == id {
			return index, true
		}
	}
	return -1, false
}
