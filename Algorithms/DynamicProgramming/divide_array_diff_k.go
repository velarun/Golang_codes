package main

import "fmt"

//Checks
func solve(array []int, size, k int) bool { 
	
	totalSum := 0
	for i:=0; i<size; i++ { 
		totalSum += array[i] 
	}

	if (totalSum - k) % 2 == 1 { 
		return false
	}
 
	S := (totalSum - k) / 2

	sum := 0
	for i:=0; i<size; i++ { 
		sum += array[i] 
		if sum == S {
			return true
		}
	}
	return false
}

func divideArray(arr []int, n int) bool { 

	sum := 0
	for i:=0; i<n; i++ { 
		sum += arr[i] 
	}

	sum_so_far := 0

	for i:=0; i<n; i++ { 
		if 2 * (sum_so_far + arr[i]) == sum { 
			fmt.Println("The array can be divided into two subarrays with equal sum")
			fmt.Println(arr[0:i+1]) 
			fmt.Println(arr[i+1:n]) 
			return true
		}
		sum_so_far += arr[i] 
	}

	fmt.Println("The array cannot be divided into two subarrays with equal sum") 

	return false
}

func main() { 
	array := []int{2, 4, 1, 5, 2} 
	k := 1
	n := len(array)
	if solve(array, n, k) { 
		fmt.Println("Yes") 
	} else { 
		fmt.Println("No") 
	}

	if divideArray(array, n) { 
		fmt.Println("Yes") 
	} else { 
		fmt.Println("No") 
	}
}