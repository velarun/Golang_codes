package main

import "fmt"

type Graphs map[int]Graph

type Graph []int

func (g Graphs) addEdge(u int, v int) {
	g[u] = append(g[u], v)
}

func (g Graphs) DFSUtil(v int, visited map[int]bool) {

	visited[v] = true
	fmt.Print(v, " ")

	for _, i := range g[v] {
		if visited[i] == false {
			g.DFSUtil(i, visited)
		}
	}
}

func (g Graphs) DFS(v int) {

	//Mark all the vertices as not visited
	visited := make(map[int]bool)
	g.DFSUtil(v, visited)
}

func main() {

	var g Graphs = map[int]Graph{}
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)

	fmt.Println("Following is DFS from (starting from vertex 2)")
	g.DFS(2)
	fmt.Println()

}
