package lemin

import (
	"errors"
)

type Farm struct {
	Rooms []*Room
}

type Room struct {
	Name    string
	Tunnel  []*Room
	Weight  map[string]int
	Visited map[string]bool
}

//AddRoom adds name of rooms to ant farm
func (f *Farm) AddRoom(name string) error { // err

	if contains(f.Rooms, name) {
		return errors.New("ERROR: two or more rooms with similar name")
	}

	newRoom := &Room{Name: name, Weight: map[string]int{}}
	f.Rooms = append(f.Rooms, newRoom)

	return nil
}

//AddTunnels adds tunnels and gives to all weight
func (f *Farm) AddTunnels(from, to string) error {
	fromRoom := f.GetRoom(from)
	toRoom := f.GetRoom(to)

	if fromRoom == nil || toRoom == nil {
		return errors.New("ERROR: creating tunnel between invalid room/rooms")
	} else if contains(fromRoom.Tunnel, to) {
		return errors.New("ERROR: creating existing tunnel")
	} else if fromRoom == toRoom {
		return errors.New("ERROR: cannot creat tunnel between similar room")
	} else {
		fromRoom.Weight[to] = 1
		toRoom.Weight[from] = 1

		fromRoom.Tunnel = append(fromRoom.Tunnel, toRoom)
		toRoom.Tunnel = append(toRoom.Tunnel, fromRoom)
	}
	return nil
}

//GetRoom is giving Room struct information by name of the Room
func (f *Farm) GetRoom(name string) *Room {
	for _, v := range f.Rooms {
		if v.Name == name {
			return v
		}
	}
	return nil
}

//Print is for Ant Farm Graph Visualisation
func (f *Farm) Print(end string) ([]string, bool) {
	var arr []string
	arr = append(arr, end)

	for x := end; p[x] != ""; x = p[x] {
		arr = append(arr, p[x])

	}
	if dist[end] == 1e8 {
		return arr, false
	}
	return arr, true
}

func contains(s []*Room, name string) bool {
	for _, v := range s {
		if name == v.Name {
			return true
		}
	}
	return false
}

//DeleteTunnel is deleting tunnels between rooms
func (f *Farm) DeleteTunnel(from, to string) {
	x := f.GetRoom(from)
	y := f.GetRoom(to)
	var rplc []*Room
	for _, v := range x.Tunnel {
		if v.Name != y.Name {
			rplc = append(rplc, v)
		}
	}
	x.Tunnel = rplc
}
