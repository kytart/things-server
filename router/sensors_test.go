package router

import (
	"fmt"
	"github.com/kytart/things-server/sensors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetLastReadingOk(t *testing.T) {
	request, err := http.NewRequest("GET", "/sensors/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	recordedAt := time.Now()
	sensors.RecordValue("1", 10, recordedAt)

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, request)

	if rr.Code != http.StatusOK {
		t.Errorf("wrong response code: expected %v got %v", rr.Code, http.StatusOK)
	}

	expectedBody := fmt.Sprintf(`{"id":"1","value":%d,"read_at":%v}`, 10, recordedAt.Unix())
	actualBody := strings.TrimSpace(rr.Body.String())
	if expectedBody != actualBody {
		t.Errorf("unexpected body: expected %v got %v", expectedBody, actualBody)
	}
}

func TestGetLastReadingNotFound(t *testing.T) {
	request, err := http.NewRequest("GET", "/sensors/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, request)

	if rr.Code != http.StatusNotFound {
		t.Errorf("wrong response code: expected %v got %v", rr.Code, http.StatusNotFound)
	}
}

func TestRecordReading(t *testing.T) {
	requestBody := strings.NewReader(`{"value":20}`)
	request, err := http.NewRequest("POST", "/sensors/3", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := CreateRouter()
	router.ServeHTTP(rr, request)

	if rr.Code != http.StatusOK {
		t.Errorf("wrong response code: expected %v got %v", rr.Code, http.StatusOK)
	}

	recordedSensor := sensors.GetSensorById("3")
	if recordedSensor == nil {
		t.Error("expected creation of new sensor but it wasn't found")
	} else if recordedSensor.Value != 20 {
		t.Errorf("wrong recorded value: expected 20 got %v", recordedSensor.Value)
	}
}
