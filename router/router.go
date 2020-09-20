package router

import "github.com/gorilla/mux"

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	HandleSensorsRequests(router)
	return router
}
