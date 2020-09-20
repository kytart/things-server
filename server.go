package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kytart/things-server/router"
)

func main() {
	const PORT = 8080
	log.Printf("Listening on localhost:%d\n", PORT)
	myRouter := router.CreateRouter()
	host := fmt.Sprintf(":%d", PORT)
	log.Fatal(http.ListenAndServe(host, myRouter))
}
