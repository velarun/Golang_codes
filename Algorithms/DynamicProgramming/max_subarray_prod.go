package main

import (
	"fmt"
	"math"
)

//O(n)
func maxSubarrayProduct(arr []int, n int) int { 
	// max positive product ending at the current position 
	max_ending_here := 1 

	// min negative product ending at the current position 
	min_ending_here := 1 

	// Initialize overall max product 
	max_so_far := 1 
	flag := 0 
	/* Traverse through the array. Following values are 
	maintained after the i'th iteration: 
	max_ending_here is always 1 or some positive product 
					ending with arr[i] 
	min_ending_here is always 1 or some negative product 
					ending with arr[i] */
	for i := 0; i < n; i++ { 
		/* If this element is positive, update max_ending_here. 
		Update min_ending_here only if min_ending_here is 
		negative */
		if (arr[i] > 0) { 
			max_ending_here = max_ending_here * arr[i]
			min_ending_here = int(math.Min(float64(min_ending_here * arr[i]), float64(1)))
			flag = 1
		} else if (arr[i] == 0) { 
			max_ending_here = 1 
			min_ending_here = 1
		} else { 
			temp := max_ending_here 
			max_ending_here = int(math.Max(float64(min_ending_here * arr[i]), float64(1))) 
			min_ending_here = temp * arr[i]
		} 

		// update max_so_far, if needed 
		if max_so_far < max_ending_here {
			max_so_far = max_ending_here 
		}
	} 
	if flag == 0 && max_so_far == 1 { 
		return 0
	}

	return max_so_far 
} 

func main() { 
	arr := []int{1, -2, -3, 0, 7, -8, -2} 
	n := len(arr) 
	fmt.Println("Maximum Sub array product is ", maxSubarrayProduct(arr, n))  
} 