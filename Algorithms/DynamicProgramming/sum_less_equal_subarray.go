package main

import (
	"fmt"
	"math"
)

func findMaxSubarraySum(arr []int, n, sum int) int { 

	curr_sum, max_sum, start := arr[0], 0, 0

	// To find max_sum less than sum 
	for i := 1; i < n; i++ { 

		// Update max_sum if it becomes 
		// greater than curr_sum 
		if curr_sum <= sum { 
			max_sum = int(math.Max(float64(max_sum), float64(curr_sum)))
		} 

		// If curr_sum becomes greater than 
		// sum subtract starting elements of array 
		for curr_sum + arr[i] > sum && start < i { 
			curr_sum -= arr[start]
			start++
		} 
		
		// Add elements to curr_sum 
		curr_sum += arr[i]
	} 

	// Adding an extra check for last subarray 
	if curr_sum <= sum {
		max_sum = int(math.Max(float64(max_sum), float64(curr_sum)))
	}

	return max_sum
}

func main() { 
	arr := []int{ 1, 2, 3, 4, 5 } 
	n := len(arr) 
	sum := 11 

	fmt.Println(findMaxSubarraySum(arr, n, sum)) 
} 
