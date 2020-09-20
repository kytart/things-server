package sensors

import "time"

type SensorsStore interface {
	GetSensorById(id string) *Sensor
	RecordValue(id string, value int, recordedAt time.Time) error
}
