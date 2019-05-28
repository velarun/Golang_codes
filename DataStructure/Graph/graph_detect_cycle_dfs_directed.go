package main

import "fmt"

type DGraphs struct {
	vertices int
	graph    Graphs
}

type Graphs map[int]Graph
type Graph []int

func (g Graphs) addEdge(u int, v int) {
	g[u] = append(g[u], v)
}

func (g DGraphs) isCyclicUtil(v int, visited, record map[int]bool) bool {

	visited[v] = true
	record[v] = true

	for _, i := range g.graph[v] {
		if visited[i] == false {
			if g.isCyclicUtil(i, visited, record) {
				return true
			}
		} else if record[i] == true {
			return true
		}
	}

	record[v] = false
	return false
}

func (g DGraphs) isCyclic() bool {

	//Mark all the vertices as not visited
	visited := make(map[int]bool)
	record := make(map[int]bool)

	for i := 0; i < g.vertices; i++ {
		if visited[i] == false {
			if g.isCyclicUtil(i, visited, record) {
				return true
			}
		}
	}

	return false
}

func main() {

	var g1 Graphs = map[int]Graph{}
	g1.addEdge(1, 0)
	g1.addEdge(0, 2)
	g1.addEdge(2, 0)
	g1.addEdge(0, 3)
	g1.addEdge(3, 4)
	g := DGraphs{vertices: 5, graph: g1}

	if g.isCyclic() {
		fmt.Println("Graph contains cycle")
	} else {
		fmt.Println("Graph does not contain cycle")
	}

	var g2 Graphs = map[int]Graph{}
	g2.addEdge(0, 1)
	g2.addEdge(1, 2)
	p := DGraphs{vertices: 3, graph: g2}

	if p.isCyclic() {
		fmt.Println("Graph contains cycle")
	} else {
		fmt.Println("Graph does not contain cycle")
	}

}
