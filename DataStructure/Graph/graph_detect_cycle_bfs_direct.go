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

func (g DGraphs) isCyclicUtil(s int, visited map[int]bool, n int) bool {

	// Set parent vertex for every vertex as -1.
	parent := make([]int, n)
	for m := 0; m < n; m++ {
		parent[m] = -1
	}

	queue := []int{}

	//Mark the source node as visited and enqueue it
	queue = append(queue, s)
	visited[s] = true

	for len(queue) > 0 {
		// Dequeue a vertex from queue and print it
		u := queue[0]
		queue = queue[1:]
		for _, v := range g.graph[u] {
			if visited[v] == false {
				queue = append(queue, v)
				visited[v] = true
				parent[v] = u
			} else if parent[u] != v {
				return true
			}
		}
	}
	return false
}

func (g DGraphs) isCyclicBFSconntected() bool {

	//Mark all the vertices as not visited
	visited := make(map[int]bool)

	for i := 0; i < g.vertices; i++ {
		if visited[i] == false {
			if g.isCyclicUtil(i, visited, g.vertices) {
				return true
			}
		}
	}

	return false
}

func main() {

	var g1 Graphs = map[int]Graph{}
	g1.addEdge(0, 1)
	g1.addEdge(1, 2)
	g1.addEdge(2, 0)
	g1.addEdge(2, 3)
	g := DGraphs{vertices: 4, graph: g1}

	if g.isCyclicBFSconntected() {
		fmt.Println("Graph contains cycle")
	} else {
		fmt.Println("Graph does not contain cycle")
	}

	var g2 Graphs = map[int]Graph{}
	g2.addEdge(0, 1)
	g2.addEdge(1, 2)
	p := DGraphs{vertices: 3, graph: g2}

	if p.isCyclicBFSconntected() {
		fmt.Println("Graph contains cycle")
	} else {
		fmt.Println("Graph does not contain cycle")
	}

}
