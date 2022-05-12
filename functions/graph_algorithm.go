package lemin

import (
	"fmt"
	"math"
	"strconv"
)

var p = make(map[string]string)
var dist = make(map[string]int)

//BFS is looking for shortest way
func (f *Farm) BFS(start string) {
	var route []string
	for _, v := range f.Rooms {
		dist[v.Name] = 1e8
	}
	dist[start] = 0
	route = append(route, start)
	for len(route) > 0 {
		b := true
		for i := 0; i < len(route)-1; i++ {
			if dist[route[i]] != dist[route[i+1]] {
				b = false
			}
		}
		if b {
			fmt.Println(route)
		}
		curr := route[0]
		route = route[1:]
		to := f.GetRoom(curr)
		for _, val := range to.Tunnel {
			length := to.Weight[val.Name]
			if dist[curr]+length < dist[val.Name] {
				dist[val.Name] = dist[curr] + length
				route = append(route, val.Name)
				p[val.Name] = curr
			}
		}
	}
}

//DeleteAdjacent is deleting one way tunnel between rooms
func (f *Farm) DeleteAdjacent(end string) {
	v := f.GetRoom(end)
	v.Weight[p[end]] = -1
	fmt.Printf("%v --> ", end)
	for x := end; p[x] != ""; x = p[x] {
		fmt.Printf("%v --> ", p[x])
		v := f.GetRoom(p[x])
		if p[p[x]] != "" {
			v.Weight[p[p[x]]] = -1
		}

		var rplc []*Room
		for _, val := range v.Tunnel {
			if val.Name != x {
				rplc = append(rplc, val)
			}
		}
		v.Tunnel = rplc
	}
	fmt.Println()
}

//CreatingFinalFarm is deleting paths with repetitive direction to avoid traffic jams
func (f *Farm) CreatingFinalFarm(paths [][]string) {
	used := make(map[string]int)

	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[i]); j++ {
			if j+1 < len(paths[i]) {
				used[paths[i][j]+paths[i][j+1]]++
				used[paths[i][j+1]+paths[i][j]]++
				if used[paths[i][j]+paths[i][j+1]] > 1 {
					f.DeleteTunnel(paths[i][j], paths[i][j+1])
					f.DeleteTunnel(paths[i][j+1], paths[i][j])
				}
			}
		}
	}
}

//OneRoomProblem is solving AntFarm with one room in the middle traffic problem
func OneRoomProblem(paths *[][]string) {
	visited := make(map[string]int)
	indexes := []int{}
	for i := 0; i < len(*paths); i++ {
		for j := 0; j < len((*paths)[i]); j++ {
			visited[(*paths)[i][j]]++
			if visited[(*paths)[i][j]] > 1 && j > 0 && j < len((*paths)[i])-1 {
				indexes = append(indexes, i)
			}
		}
	}

	for i := 0; i < len(indexes); i++ {
		*paths = append((*paths)[:indexes[i]], (*paths)[indexes[i]+1:]...)
	}
}

//PrintResult is printing how ants are moving step by step
func PrintResult(paths [][]string, ant int, endRoom string) {
	sizes := make([]int, len(paths))
	sum := 0
	exception := false
	for i := 0; i < len(paths); i++ {
		paths[i] = paths[i][:len(paths[i])-1]
		if len(paths[i]) == 1 {
			exception = true
		}
		sizes[i] = len(paths[i])
		sum += sizes[i]
	}

	n := 1

	m := (ant - (len(paths)*len(paths[len(paths)-1]) - sum))

	m = int(math.Ceil(float64(m) / float64(len(paths))))
	m += len(paths[len(paths)-1])

	if exception {
		m = 2
	}
	ans := make([][]string, m-1)
	fmt.Println(paths, "\n\n")
LOOP:
	for i := 0; ; i++ {
		for j := 0; j < len(paths); j++ {
			if n-1 == ant {
				break LOOP
			}
			if i >= len(paths[j]) {
				cnt := len(paths[j]) - sizes[j]
				if len(paths[j]) == 1 {
					for k := 1; k <= ant; k++ {
						ans[0] = append(ans[0], "L"+strconv.Itoa(k)+"-"+endRoom)
					}
					break LOOP
				}
				for k := len(paths[j]) - cnt - 1; k >= 0; k-- {
					ans[cnt] = append(ans[cnt], "L"+strconv.Itoa(n)+"-"+paths[j][k])
					cnt++
				}

				paths[j] = append(paths[j], "_"+strconv.Itoa(n))
				n++
			}
		}
	}

	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(ans[i]); j++ {
			fmt.Print(ans[i][j], " ")
		}
		fmt.Println()
	}

}
