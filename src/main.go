package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = 8080

var rooms = []Room{
	Room{Id: "1", Name: "living_room", Temperature: 20},
	Room{Id: "2", Name: "bedroom", Temperature: 22},
}

type Room struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Temperature int    `json:"temperature"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func getRooms(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(rooms)
}

func getRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, room := range rooms {
		if room.Id == key {
			json.NewEncoder(w).Encode(room)
		}
	}
}

func addRoom(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var room Room
	json.Unmarshal(body, &room)
	rooms = append(rooms, room)
	json.NewEncoder(w).Encode(room)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/rooms", getRooms)
	router.HandleFunc("/room/{id}", getRoom)
	router.HandleFunc("/room", addRoom).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), router))
}

func main() {
	log.Println("Listening on localhost:%d", PORT)
	handleRequests()
}
