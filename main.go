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

	g, s, e := GraphConstruct(text)

	//	step 1: dijkstra
	g.BFS(s)
	//	step 2: inversing edges from found path with negative costs
	g.DeleteAdjacent(e)

	// step 2.1: duplicate all intermediate vertices

	g.Print(e)
}
