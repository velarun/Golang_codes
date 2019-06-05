package main

import (
	"fmt"
	"math"
)

func areConsecutives(arr []int, n int) bool {

	first_term := math.MaxInt64
	arr_sum := 0
	for i:=0; i<n; i++ {
		arr_sum += arr[i]
		if arr[i] < first_term { 
			first_term = arr[i] 
		}
	}
	
	ap_sum := (n * (2 * first_term + (n - 1) * 1)) / 2 
	
	return ap_sum == arr_sum 
}

func main() {
	arr := []int{2, 1, 0, -3, -1, -2} 
	n := len(arr) 
	if areConsecutives(arr, n) { 
		fmt.Println("Yes") 
	} else { 
		fmt.Println("No")
	}
}

