package router

import (
	"encoding/json"
	"github.com/kytart/things-server/sensors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func createGetLastReadingHandler(sensorsStore sensors.SensorsStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		sensor := sensorsStore.GetSensorById(id)
		if sensor == nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			err := json.NewEncoder(w).Encode(sensor)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

func createRecordReadingHandler(sensorsStore sensors.SensorsStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		body, _ := ioutil.ReadAll(r.Body)
		var reading struct {
			Value int `json:"value"`
		}
		err := json.Unmarshal(body, &reading)
		if err == nil {
			recordedAt := time.Now()
			err = sensorsStore.RecordValue(id, reading.Value, recordedAt)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func HandleSensorsRequests(
	router *mux.Router,
	sensorsStore sensors.SensorsStore,
) {
	router.HandleFunc("/sensors/{id}", createGetLastReadingHandler(sensorsStore)).Methods("GET")
	router.HandleFunc("/sensors/{id}", createRecordReadingHandler(sensorsStore)).Methods("POST")
}
