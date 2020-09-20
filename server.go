package main

import (
	"fmt"
	"github.com/kytart/things-server/sensors"
	"log"
	"net/http"

	"github.com/kytart/things-server/router"
)

func main() {
	const PORT = 8080
	log.Printf("Listening on localhost:%d\n", PORT)
	sensorsStore := sensors.NewInMemorySensorsStore([]sensors.Sensor{})
	myRouter := router.CreateRouter(sensorsStore)
	host := fmt.Sprintf(":%d", PORT)
	log.Fatal(http.ListenAndServe(host, myRouter))
}
