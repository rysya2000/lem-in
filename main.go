package main

import (
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

	antFarm, startRoom, endRoom, ant, err := FileProcessing(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	antFarmReversed := antFarm
	antFarmReversed.BFS(startRoom)

	var paths [][]string
	arr, ok := antFarmReversed.Print(endRoom)
	if ok {
		paths = append(paths, arr)
	}
	v := antFarmReversed.GetRoom(startRoom)

	startroomTunnels := 0
	for range v.Tunnel {
		startroomTunnels++
	}
	for startroomTunnels > 0 {
		antFarmReversed.DeleteAdjacent(endRoom)

		if startroomTunnels != 1 {
			antFarmReversed.BFS(startRoom)
			arr, ok := antFarmReversed.Print(endRoom)
			if ok {
				paths = append(paths, arr)
			}
		}
		startroomTunnels--
	}

	res, _, _, _, _ := FileProcessing(text)
	res.CreatingFinalFarm(paths)

	paths = [][]string{}
	for {
		res.BFS(startRoom)
		res.DeleteAdjacent(endRoom)
		arr, ok := res.Print(endRoom)
		if !ok {
			break
		}
		paths = append(paths, arr)
	}

	PrintResult(paths, ant, endRoom)
}
