package router

import (
	"encoding/json"
	"github.com/kytart/things-server/sensors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getLastReading(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	sensor := sensors.GetSensorById(id)
	if sensor == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		err := json.NewEncoder(w).Encode(sensor)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func recordReading(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	body, _ := ioutil.ReadAll(r.Body)
	var reading struct {
		Value int `json:"value"`
	}
	err := json.Unmarshal(body, &reading)
	if err == nil {
		readAt := time.Now()
		sensors.RecordValue(id, reading.Value, readAt)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func HandleSensorsRequests(router *mux.Router) {
	router.HandleFunc("/sensors/{id}", getLastReading).Methods("GET")
	router.HandleFunc("/sensors/{id}", recordReading).Methods("POST")
}
