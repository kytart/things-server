package router

import (
	"github.com/gorilla/mux"
	"github.com/kytart/things-server/sensors"
)

func CreateRouter(sensorsStore sensors.SensorsStore) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	HandleSensorsRequests(router, sensorsStore)
	return router
}
