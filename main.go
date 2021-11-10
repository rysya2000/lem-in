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

	antFarm, startRoom, endRoom, ant := FileProcessing(text)

	fmt.Println(ant)
	//	step 1: dijkstra
	antFarm.BFS(startRoom)

	antFarm.Print(endRoom)

	//	step 2: inversing edges from found path with negative costs
	//!!!! have problem with negative costs in example00.txt
	antFarm.DeleteAdjacent(endRoom)

	antFarm.FordBellman(startRoom)

	// step 2.1: duplicate all intermediate vertices
	antFarm.Print(endRoom)
}
