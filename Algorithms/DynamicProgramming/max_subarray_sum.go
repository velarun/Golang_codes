package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func max_sum(a []int, size int) int {
	max_sum := a[0]
	curr_max := a[0]

	for i := 1; i < size; i++ {
		curr_max = max(a[i], curr_max+a[i])
		max_sum = max(curr_max, max_sum)
	}

	return max_sum
}

func main() {
	var a []int = []int{-1, -2, 4, 1, -2, -7, 4}
	if len(a) > 1 {
		fmt.Println(max_sum(a, len(a)))
	} else {
		fmt.Println("List should be greater than 1")
	}
}
