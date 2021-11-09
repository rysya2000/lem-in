package lemin

import (
	"fmt"
	"strings"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	Name     string
	adjacent []*Vertex
	weights  map[string]int
}

func (g *Graph) AddVertex(name string) {
	if contains(g.vertices, name) {
		fmt.Printf("vertex %v already exiting", name)
	} else {
		newVertex := &Vertex{Name: name, weights: map[string]int{}}
		g.vertices = append(g.vertices, newVertex)
	}
}

func (g *Graph) AddEdge(from, to string) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v --> %v)", fromVertex, toVertex)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v --> %v)", from, to)
		fmt.Println(err)
	} else {
		fromVertex.weights[to] = 1
		toVertex.weights[from] = 1

		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		toVertex.adjacent = append(toVertex.adjacent, fromVertex)
	}
}

func (g *Graph) getVertex(name string) *Vertex {
	for _, v := range g.vertices {
		if v.Name == name {
			return v
		}
	}
	return nil
}

func contains(s []*Vertex, name string) bool {
	for _, v := range s {
		if name == v.Name {
			return true
		}
	}
	return false
}

func (g *Graph) Print(end string) {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v (%v) : ", v.Name, dist[v.Name])
		for _, v2 := range v.adjacent {
			fmt.Printf("[%v %v] ", v2.Name, v.weights[v2.Name])
		}
	}
	fmt.Print("\n\n", end)
	for x := end; p[x] != ""; x = p[x] {
		fmt.Print(" --> ", p[x])
	}
	fmt.Println()
}

func GraphConstruct(text []string) (Graph, string, string) {
	test := &Graph{}
	start := ""
	end := ""
	dot := ""
	for i, t := range text {
		if i == 0 {
			continue
		}
		if t == "##start" {
			dot = "s"
			continue
		}
		if t == "##end" {
			dot = "e"
			continue
		}
		if strings.Contains(t, "-") {
			r := strings.Split(t, "-")
			test.AddEdge(r[0], r[1])
			continue
		}
		s := strings.Split(t, " ")
		if dot == "s" {
			start = s[0]
			dot = ""
		}
		if dot == "e" {
			end = s[0]
			dot = ""
		}
		test.AddVertex(s[0])
	}
	return *test, start, end
}

var p = make(map[string]string)
var dist = make(map[string]int)

// dijkstra algorithm
func (g *Graph) BFS(start string) {
	var Q []string
	for _, v := range g.vertices {
		dist[v.Name] = 1e8
	}
	dist[start] = 0
	Q = append(Q, start)
	for len(Q) > 0 {
		//		fmt.Println(Q)
		curr := Q[0]
		Q = Q[1:]
		to := g.getVertex(curr)
		for _, val := range to.adjacent {
			length := to.weights[val.Name]
			if dist[curr]+length < dist[val.Name] {
				dist[val.Name] = dist[curr] + length
				Q = append(Q, val.Name)
				p[val.Name] = curr
			}
		}

	}
}

func (g *Graph) DeleteAdjacent(end string) {
	v := g.getVertex(end)
	v.weights[p[end]] = -1
	for x := end; p[x] != ""; x = p[x] {
		v := g.getVertex(p[x])
		if p[p[x]] != "" {
			v.weights[p[p[x]]] = -1
		}
		var rplc []*Vertex
		for _, val := range v.adjacent {
			if val.Name != x {
				rplc = append(rplc, val)
			}
		}
		v.adjacent = rplc
	}
}

func (g *Graph) Duplicate() {

}
