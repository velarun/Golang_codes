package main

import "fmt"

type Graphs map[int]Graph
type Graph []int

func (g Graphs) addEdge(u int, v int) {
	g[u] = append(g[u], v)
	g[v] = append(g[v], u)
}

func (g Graphs) greedyColoring(leng int) {

	result := make([]int, leng)
	for m := 0; m < leng; m++ {
		result[m] = -1
	}
	result[0] = 0

	available := make([]bool, leng)

	for u := 1; u < leng; u++ {
		for _, i := range g[u] {
			if result[i] != -1 {
				available[result[i]] = true
			}
		}

		var cr int
		for cr = 0; cr < leng; cr++ {
			if available[cr] == false {
				break
			}
		}

		result[u] = cr

		for m := 0; m < leng; m++ {
			if result[m] != -1 {
				available[m] = false
			}
		}
	}

	for u := 0; u < leng; u++ {
		fmt.Println("Vertex ", u, " ---> Color ", result[u])
	}
}

func main() {

	var g1len int = 5
	var g1 Graphs = map[int]Graph{}
	g1.addEdge(0, 1)
	g1.addEdge(0, 2)
	g1.addEdge(1, 2)
	g1.addEdge(1, 3)
	g1.addEdge(2, 3)
	g1.addEdge(3, 4)
	fmt.Println("Coloring of graph 1:")
	g1.greedyColoring(g1len)

	fmt.Println()

	var g2 Graphs = map[int]Graph{}
	g2.addEdge(0, 1)
	g2.addEdge(0, 2)
	g2.addEdge(1, 2)
	g2.addEdge(1, 4)
	g2.addEdge(2, 4)
	g2.addEdge(4, 3)
	fmt.Println("Coloring of graph 2:")
	g2.greedyColoring(g1len)

}
