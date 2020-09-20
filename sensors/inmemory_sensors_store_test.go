package sensors

import (
	"testing"
	"time"
)

func TestInMemorySensorsStore_GetSensorById_ReturnsNil(t *testing.T) {
	sensorsStore := NewInMemorySensorsStore([]Sensor{})
	returnedSensor := sensorsStore.GetSensorById("1")
	if returnedSensor != nil {
		t.Errorf("expected nil but got %v", returnedSensor)
	}
}

func TestInMemorySensorsStore_GetSensorById_ReturnsSensor(t *testing.T) {
	storedSensor := Sensor{
		Id:         "1",
		Value:      10,
		RecordedAt: time.Now(),
	}
	sensorsStore := NewInMemorySensorsStore([]Sensor{storedSensor})
	returnedSensor := sensorsStore.GetSensorById("1")
	if returnedSensor.Id != storedSensor.Id ||
		returnedSensor.Value != storedSensor.Value ||
		!returnedSensor.RecordedAt.Equal(storedSensor.RecordedAt) {
		t.Errorf("expected %v but got %v", storedSensor, returnedSensor)
	}
}

func TestInMemorySensorsStore_RecordValue(t *testing.T) {
	sensorsStore := NewInMemorySensorsStore([]Sensor{})
	recordedAt := time.Now()
	err := sensorsStore.RecordValue("1", 20, recordedAt)
	if err != nil {
		t.Fatal(err)
	}
	if len(sensorsStore.Sensors) != 1 {
		t.Errorf("expected 1 sensor to be in store, got %v", len(sensorsStore.Sensors))
	}
	storedSensor := sensorsStore.Sensors[0]
	if storedSensor.Id != "1" ||
		storedSensor.Value != 20 ||
		!storedSensor.RecordedAt.Equal(recordedAt) {
		expectedSensor := Sensor{
			Id:         "1",
			Value:      20,
			RecordedAt: recordedAt,
		}
		t.Errorf("expected %v but got %v", expectedSensor, storedSensor)
	}
}
