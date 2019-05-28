package main

import "fmt"

const MAX_SIZE = 9223372036854775807

type Graphs []Graph

type Graph []int

func (g Graphs) minDistance(dist []int, sptSet []bool) int {

	min := MAX_SIZE
	vertex := len(g)
	var min_index int

	for v := 0; v < vertex; v++ {
		if dist[v] < min && sptSet[v] == false {
			min = dist[v]
			min_index = v
		}
	}
	return min_index
}

func (g Graphs) dijkstra(src, dest int) {

	vertex := len(g)
	dist := make([]int, vertex)
	for i := range dist {
		dist[i] = MAX_SIZE
	}

	dist[src] = 0
	sptSet := make([]bool, vertex)

	for cout := 0; cout < vertex; cout++ {
		u := g.minDistance(dist, sptSet)
		sptSet[u] = true

		for v := 0; v < vertex; v++ {
			if (g[u][v] > 0) && (sptSet[v] == false) && (dist[v] > dist[u]+g[u][v]) {
				dist[v] = dist[u] + g[u][v]
			}
		}
	}

	fmt.Println("from", src, "to", dest, "is", dist[dest])
}

func main() {

	var p Graphs
	var q Graph
	p = append(p, append(q, []int{0, 4, 0, 0, 0, 0, 0, 8, 0}...))
	p = append(p, append(q, []int{4, 0, 8, 0, 0, 0, 0, 11, 0}...))
	p = append(p, append(q, []int{0, 8, 0, 7, 0, 4, 0, 0, 2}...))
	p = append(p, append(q, []int{0, 0, 7, 0, 9, 14, 0, 0, 0}...))
	p = append(p, append(q, []int{0, 0, 0, 9, 0, 10, 0, 0, 0}...))
	p = append(p, append(q, []int{0, 0, 4, 14, 10, 0, 2, 0, 0}...))
	p = append(p, append(q, []int{0, 0, 0, 0, 0, 2, 0, 1, 6}...))
	p = append(p, append(q, []int{8, 11, 0, 0, 0, 0, 1, 0, 7}...))
	p = append(p, append(q, []int{0, 0, 2, 0, 0, 0, 6, 7, 0}...))

	p.dijkstra(0, 8)
}
