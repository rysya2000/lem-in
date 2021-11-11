package main

import (
	"container/list"
	"fmt"
	"os"

	. "lemin/functions"
)

func main() {

	text, err := OpenFile(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	antFarm, startRoom, endRoom, ant := FileProcessing(text)

	fmt.Println(ant)
	//	step 1: dijkstra
	antFarm.BFS(startRoom)
	antFarm.Print(endRoom)

	antFarmReversed := &antFarm
	v := antFarmReversed.GetRoom(startRoom)
	cnt := 0
	for _, _ = range v.Tunnel {
		cnt++
	}
	L := list.New()
	for cnt > 0 {
		// step 3: dijkstra
		antFarmReversed.DeleteAdjacent(endRoom)

		//	step 2: inversing edges from found path with negative costs
		if cnt != 1 {
			antFarmReversed.BFS(startRoom)
			arr := antFarmReversed.Print(endRoom)
			L.PushBackList(&arr)
		}
		cnt--
	}
	for _, v := range L.Len() {
		for _, v2 := range v {

		}
	}
}
