package main

import "fmt"


var max int = 9999999999

//O(n^m)
func minCost(cost [][]int, m, n int) int {
	if n < 0 || m < 0 {
		return max
	} else if m == 0 && n == 0 {
		return cost[m][n]
	} else {
		return cost[m][n] + min(minCost(cost, m-1, n-1), minCost(cost, m-1, n), minCost(cost, m, n-1))
	}
}

//O(nm)
func minCostTimeComplexity(cost [][]int, m, n, R, C int) int {

	tc := make([][]int, C)
	for x := 0; x < R; x++ {
		tc[x] = make([]int, C)
	}

	tc[0][0] = cost[0][0]

	for i := 1; i < m+1; i++ {
		tc[i][0] = tc[i-1][0] + cost[i][0]
	}

	for j := 1; j < n+1; j++ {
		tc[0][j] = tc[0][j-1] + cost[0][j]
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			tc[i][j] = min(tc[i-1][j-1], tc[i-1][j], tc[i][j-1]) + cost[i][j]
		}
	}

	return tc[m][n]
}

func min(x, y, z int) int {
	if x < y {
		if x < z {
			return x
		} else {
			return z
		}
	} else {
		if y < z {
			return y
		} else {
			return z
		}
	}
}

func main() {
	R := 3
	C := 3
	cost := [][]int{[]int{1, 2, 3}, []int{4, 1, 2}, []int{1, 5, 3}}
	fmt.Println(minCost(cost, 2, 2))
	fmt.Println(minCostTimeComplexity(cost, 2, 2, R, C))
}
