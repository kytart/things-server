package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

const PORT = 8080

var rooms = []Room{
	Room{Name: "living_room", Temperature: 20},
	Room{Name: "bedroom", Temperature: 22},
}

type Room struct {
	Name        string `json:"name"`
	Temperature int    `json:"temperature"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func getRooms(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(rooms)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/rooms", getRooms)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}

func main() {
	log.Println("Listening on localhost:%d", PORT)
	handleRequests()
}
