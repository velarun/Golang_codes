package main

import "fmt"

//O(n) return index or sum
func midSum(a []int, n int) int {

	prefix := make([]int, n)
	suffix := make([]int, n)

	prefix[0] = a[0]
	for i:=1; i<n; i++ {
		prefix[i] = prefix[i-1] + a[i]
	}

	suffix[n-1] = a[n-1]
	for i:=n-2; i>-1; i-- {
		suffix[i] = suffix[i+1] + a[i]
	}

	for i:=0; i<n; i++ {
		if prefix[i] == suffix[i] {
			return i
		}
	}

	return -1
}

func midSumPointer(arr []int, n int) int {

	right_sum, left_sum := 0, 0
  
    //Computing right_sum 
    for i:=1; i<n; i++ { 
		right_sum += arr[i] 
	}
  
    for i, j := 0, 1; j < n; i, j = i+1, j+1 { 
        right_sum -= arr[j] 
        left_sum += arr[i] 
  
        if left_sum == right_sum { 
			return arr[i + 1] 
		}
	}

	return -1
}

func main() {
	a := []int{1, 4, 2, 5}
	n := len(a)

	fmt.Println(midSum(a, n))
	fmt.Println(midSumPointer(a, n))
}