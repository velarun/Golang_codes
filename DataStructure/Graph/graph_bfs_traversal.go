package main

import "fmt"

type Graphs map[int]Graph

type Graph []int

func (g Graphs) addEdge(u int, v int) {
	g[u] = append(g[u], v)
}

func (g Graphs) BFS(s int) {

	visited := make(map[int]bool)

	//Create a queue for BFS
	queue := []int{}

	//Mark the source node as visited and enqueue it
	queue = append(queue, s)
	visited[s] = true

	for len(queue) > 0 {

		// Dequeue a vertex from queue and print it
		s = queue[0]
		queue = queue[1:]

		fmt.Print(s, " ")

		//Get all adjacent vertices of the dequeued vertex s. If a adjacent
		//has not been visited, then mark it visited and enqueue it
		for _, i := range g[s] {
			if visited[i] == false {
				queue = append(queue, i)
				visited[i] = true
			}
		}
	}
}

func main() {

	var g Graphs = map[int]Graph{}
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(2, 3)
	g.addEdge(3, 3)

	fmt.Println("Following is BFS from (starting from vertex 2)")
	g.BFS(2)
	fmt.Println()

}
