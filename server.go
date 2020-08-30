package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kytart/things-server/router"
)

func main() {
	const PORT = 8080
	log.Println("Listening on localhost:%d", PORT)
	router := router.CreateRouter()
	host := fmt.Sprintf(":%d", PORT)
	log.Fatal(http.ListenAndServe(host, router))
}
