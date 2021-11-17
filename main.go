package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	. "lemin/functions"
)

func main() {

	text, err := OpenFile(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	var paths [][]string

	antFarm, startRoom, endRoom, ant := FileProcessing(text)
	antFarmReversed := antFarm

	fmt.Println(ant)
	//	step 1: dijkstra
	antFarmReversed.BFS(startRoom)
	arr, ok := antFarmReversed.Print(endRoom)
	if ok {
		paths = append(paths, arr)
	}
	v := antFarmReversed.GetRoom(startRoom)
	cnt := 0
	for range v.Tunnel {
		cnt++
	}

	for cnt > 0 {
		// step 3: dijkstra
		antFarmReversed.DeleteAdjacent(endRoom)

		//	step 2: inversing edges from found path with negative costs
		if cnt != 1 {
			antFarmReversed.BFS(startRoom)
			arr, ok := antFarmReversed.Print(endRoom)
			if ok {
				paths = append(paths, arr)
			}

		}
		cnt--
	}

	res, _, _, _ := FileProcessing(text)
	fmt.Println(paths)
	used := make(map[string]int)

	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[i]); j++ {
			if j+1 < len(paths[i]) {
				used[paths[i][j]+paths[i][j+1]]++
				used[paths[i][j+1]+paths[i][j]]++
				if used[paths[i][j]+paths[i][j+1]] > 1 {
					fmt.Println(paths[i][j], paths[i][j+1])
					res.DeleteTunnel(paths[i][j], paths[i][j+1])
					res.DeleteTunnel(paths[i][j+1], paths[i][j])
				}
			}
		}
	}
	fmt.Println("-----------------------------------------------------")
	//	res.DFS(startRoom, endRoom, []string{startRoom}, map[string]bool{startRoom: true})

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
	fmt.Println("---------------------------------------------------------------")

	sizes := make([]int, len(paths))
	sum := 0
	for i := 0; i < len(paths); i++ {
		paths[i] = paths[i][:len(paths[i])-1]
		sizes[i] = len(paths[i])
		sum += sizes[i]
	}

	fmt.Println(paths)
	n := 1

	m := (ant - (len(paths)*len(paths[len(paths)-1]) - sum))
	// if m%len(paths) != 0 {
	// 	m /= len(paths)
	// 	m++
	// } else {
	// 	m /= len(paths)
	// }
	m = int(math.Ceil(float64(m) / float64(len(paths))))
	m += len(paths[len(paths)-1])
	fmt.Println(m)

	ans := make([][]string, m-1)

Loop:
	for i := 0; ; i++ {
		for j := 0; j < len(paths); j++ {
			if n-1 == ant {
				break Loop
			}
			if i >= len(paths[j]) {
				cnt := len(paths[j]) - sizes[j]
				if len(paths[j]) == 1 {
					for k := 1; k <= ant; k++ {
						ans[0] = append(ans[0], "L"+strconv.Itoa(k)+"-"+endRoom)
					}
					break Loop
				}
				fmt.Printf("[%v]", cnt)
				for k := len(paths[j]) - cnt - 1; k >= 0; k-- {
					ans[cnt] = append(ans[cnt], "L"+strconv.Itoa(n)+"-"+paths[j][k])
					cnt++
				}

				paths[j] = append(paths[j], "_"+strconv.Itoa(n))
				n++
			}
			fmt.Print(paths[j][i], " ")
		}
		fmt.Println()
	}

	fmt.Println("\n--------------------------------------\n")

	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(ans[i]); j++ {
			fmt.Print(ans[i][j], " ")
		}
		fmt.Println()
	}

}
