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
}

func (g *Graph) AddVertex(name string) {
	if contains(g.vertices, name) {
		fmt.Printf("vertex %v already exiting", name)
	} else {
		g.vertices = append(g.vertices, &Vertex{Name: name})
	}
}

func (g *Graph) AddEdge(from, to string) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v --> %v)", from, to)
		fmt.Println(err)
	} else {
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

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.Name)
		for _, v := range v.adjacent {
			fmt.Printf("%v ", v.Name)
		}
	}
}

func GraphConstruct(text []string) {
	test := &Graph{}
	for i, t := range text {
		if i == 0 {
			continue
		}
		if t[0] == '#' {
			continue
		}
		if strings.Contains(t, "-") {
			r := strings.Split(t, "-")
			test.AddEdge(r[0], r[1])
			continue
		}
		s := strings.Split(t, " ")
		test.AddVertex(s[0])

	}
	test.Print()
}
