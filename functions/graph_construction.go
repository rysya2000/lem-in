package lemin

import (
	"fmt"
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
func (f *Farm) AddRoom(name string) {

	if contains(f.Rooms, name) {
		fmt.Printf("Room %v already exiting", name)
	} else {
		newRoom := &Room{Name: name, Weight: map[string]int{}}
		f.Rooms = append(f.Rooms, newRoom)
	}

}

//AddTunnels adds tunnels and gives to all weight
func (f *Farm) AddTunnels(from, to string) {
	fromRoom := f.GetRoom(from)
	toRoom := f.GetRoom(to)

	if fromRoom == nil || toRoom == nil {
		err := fmt.Errorf("Invalid Room (%v --> %v)", fromRoom, toRoom)
		fmt.Println(err.Error())
	} else if contains(fromRoom.Tunnel, to) {
		err := fmt.Errorf("Existing Tunnel (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		fromRoom.Weight[to] = 1
		toRoom.Weight[from] = 1

		fromRoom.Tunnel = append(fromRoom.Tunnel, toRoom)
		toRoom.Tunnel = append(toRoom.Tunnel, fromRoom)

	}
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
	// for _, v := range f.Rooms {
	// 	fmt.Printf("\nRoom %v (%v) : ", v.Name, dist[v.Name])
	// 	for _, v2 := range v.Tunnel {
	// 		fmt.Printf("[%v %v] ", v2.Name, v.Weight[v2.Name])
	// 	}
	// }
	var arr []string
	arr = append(arr, end)

	// fmt.Print("\n\n", end)
	for x := end; p[x] != ""; x = p[x] {
		arr = append(arr, p[x])
		// fmt.Print(" --> ", p[x])
	}
	fmt.Println()
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
