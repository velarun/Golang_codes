package main

import "fmt"

type cell struct {
	x int
	y int
	dist int
}

func isInside(x, y, N int) bool {
	if x >= 1 && x <= N && y >= 1 && y <= N {
		return true
	}
	return false
}
	 
func minStepToReachTarget(knightpos, targetpos []int, N int) { 

	dx := []int{2, 2, -2, -2, 1, 1, -1, -1} 
	dy := []int{1, -1, 1, -1, 2, -2, 2, -2}
	
	var queue []cell 
	d := cell{x: knightpos[0], y: knightpos[1], dist: 0}
	queue = append(queue, d) 
	
	visited := make([][]bool, N)
	for i:=0; i<N; i++ {
		visited[i] = make([]bool, N)
	}

	visited[knightpos[0]][knightpos[1]] = true
	
	for len(queue) > 0 { 
		fmt.Println(queue)
		t := queue[0] 
		queue = queue[1:] 
		
		if t.x == targetpos[0] && t.y == targetpos[1] { 
			fmt.Println(t.dist)
		}
			 
		for i:=0; i<8; i++ { 
			x := t.x + dx[i] 
			y := t.y + dy[i] 
			
			if isInside(x, y, N) && !visited[x][y] { 
				visited[x][y] = true
				d := cell{x: knightpos[0], y: knightpos[1], dist: t.dist + 1}
				queue = append(queue, d) 
			}
		}
	}
}

func main() {
	N := 30
	knightpos := []int{1, 1} 
	targetpos := []int{30, 30} 
	minStepToReachTarget(knightpos, targetpos, N) 
}
