package lemin

import "fmt"

var p = make(map[string]string)
var dist = make(map[string]int)

func (f *Farm) BFS(start string) {
	var Q []string
	for _, v := range f.Rooms {
		dist[v.Name] = 1e8
	}
	dist[start] = 0
	Q = append(Q, start)
	for len(Q) > 0 {
		// b := true
		// for i := 0; i < len(Q)-1; i++ {
		// 	if dist[Q[i]] != dist[Q[i+1]] {
		// 		b = false
		// 	}
		// }
		// if b {
//		fmt.Println(Q)
		// }
		curr := Q[0]
		Q = Q[1:]
		to := f.GetRoom(curr)
		for _, val := range to.Tunnel {
			length := to.Weight[val.Name]
			if dist[curr]+length < dist[val.Name] {
				dist[val.Name] = dist[curr] + length
				Q = append(Q, val.Name)
				p[val.Name] = curr
			}
		}
	}
}

func (f *Farm) DeleteAdjacent(end string) {
	v := f.GetRoom(end)
	v.Weight[p[end]] = -1
	for x := end; p[x] != ""; x = p[x] {
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
}

func (f *Farm) DFS(v string, end string, path []string, used map[string]bool) {
	if v == end {
		fmt.Println(path)
		return
	}
	to := f.GetRoom(v)
	for _, val := range to.Tunnel {
		if !used[val.Name] {
			path = append(path, val.Name)
			used[val.Name] = true
			f.DFS(val.Name, end, path, used)
			path = path[:len(path)-1]
			used[val.Name] = false
		}
	}
}
