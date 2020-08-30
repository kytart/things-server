package rooms

var rooms = []Room{
	Room{Id: "1", Name: "living_room", Temperature: 20},
	Room{Id: "2", Name: "bedroom", Temperature: 22},
}

type Room struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Temperature int    `json:"temperature"`
}

func GetAllRooms() []Room {
	return rooms
}

func GetRoomById(id string) (*Room, bool) {
	for _, room := range rooms {
		if room.Id == id {
			return &room, true
		}
	}
	return nil, false
}

func AddRoom(room Room) {
	rooms = append(rooms, room)
}

func UpdateRoom(updateRoom Room) bool {
	for index, room := range rooms {
		if room.Id == updateRoom.Id {
			rooms[index] = updateRoom
			return true
		}
	}
	return false
}

func DeleteRoom(id string) bool {
	for index, room := range rooms {
		if room.Id == id {
			rooms = append(rooms[:index], rooms[index+1:]...)
			return true
		}
	}
	return false
}
