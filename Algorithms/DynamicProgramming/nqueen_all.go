package main

import (
	"fmt"
	"math"
)

var arr []int

func nQueens(k, n int) {
	for i := 1; i <= n; i++ {
		if canPlace(k, i) {
			arr[k] = i
			if k == n {
				display(n)
			} else {
				nQueens(k+1, n)
			}
		}
	}
}

func canPlace(k, i int) bool {
	for j := 1; j <= k-1; j++ {
		if arr[j] == i || (math.Abs(float64(arr[j]-i)) == math.Abs(float64(j-k))) {
			return false
		}
	}
	return true
}

func display(n int) {

	fmt.Println("Arrangement:")

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if arr[i] != j {
				fmt.Print("\t_")
			} else {
				fmt.Print("\tQ")
			}
		}
		fmt.Println("")
	}
}

func main() {
	n := 8
	arr = make([]int, n+1)

	nQueens(1, n)
}
