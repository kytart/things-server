package sensors

import "time"

type InMemorySensorsStore struct {
	Sensors []Sensor
}

func NewInMemorySensorsStore(sensors []Sensor) *InMemorySensorsStore {
	return &InMemorySensorsStore{
		Sensors: sensors,
	}
}

func (sensorsStore *InMemorySensorsStore) GetSensorById(id string) *Sensor {
	for _, sensor := range sensorsStore.Sensors {
		if sensor.Id == id {
			return &sensor
		}
	}
	return nil
}

func (sensorsStore *InMemorySensorsStore) RecordValue(id string, value int, recordedAt time.Time) error {
	index, ok := findSensorIndexById(sensorsStore.Sensors, id)
	if ok {
		sensorsStore.Sensors[index].Value = value
		sensorsStore.Sensors[index].RecordedAt = recordedAt
	} else {
		newSensor := Sensor{
			Id:         id,
			Value:      value,
			RecordedAt: recordedAt,
		}
		sensorsStore.Sensors = append(sensorsStore.Sensors, newSensor)
	}
	return nil
}

func findSensorIndexById(sensors []Sensor, id string) (int, bool) {
	for index, sensor := range sensors {
		if sensor.Id == id {
			return index, true
		}
	}
	return -1, false
}
