package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kytart/things-server/rooms"
)

func getRooms(w http.ResponseWriter, r *http.Request) {
	allRooms := rooms.GetAllRooms()
	json.NewEncoder(w).Encode(allRooms)
}

func getRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	room, ok := rooms.GetRoomById(id)
	if ok {
		json.NewEncoder(w).Encode(room)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func addRoom(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var room rooms.Room
	json.Unmarshal(body, &room)
	rooms.AddRoom(room)
	json.NewEncoder(w).Encode(room)
}

func updateRoom(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var room rooms.Room
	json.Unmarshal(body, &room)
	ok := rooms.UpdateRoom(room)
	if ok {
		json.NewEncoder(w).Encode(room)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func deleteRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ok := rooms.DeleteRoom(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}
}

func HandleRoomsRequests(router *mux.Router) {
	router.HandleFunc("/rooms", getRooms)
	router.HandleFunc("/room/{id}", getRoom).Methods("GET")
	router.HandleFunc("/room", addRoom).Methods("POST")
	router.HandleFunc("/room/{id}", updateRoom).Methods("PUT")
	router.HandleFunc("/room/{id}", deleteRoom).Methods("DELETE")
}
