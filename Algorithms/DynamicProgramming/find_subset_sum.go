package main

import "fmt"

//Negative numbers also works
func isSubsetSum(set []int, n, sum int) bool {
	if sum == 0 {
		return true
	}

	if n == 0 && sum != 0 {
		return false
	}

	if set[n-1] > sum {
		return isSubsetSum(set, n-1, sum)
	}

	return isSubsetSum(set, n-1, sum) || isSubsetSum(set, n-1, sum-set[n-1])
}

//Negative numbers also works
func subArraySum(arr []int, n, Sum int) {

	Map := make(map[int]int)
	curr_sum := 0

	for i := 0; i < n; i++ {

		curr_sum = curr_sum + arr[i]

		if curr_sum == Sum {
			fmt.Println("Sum found between indexes 0 to", i)
			return
		}

		if _, ok := Map[curr_sum-Sum]; ok {
			fmt.Println("Sum found between indexes", Map[curr_sum-Sum]+1, "to", i)
			return
		}

		Map[curr_sum] = i
	}

	print("No subarray with given sum exists")
}

func main() {
	set := []int{15, 2, 4, 8, 9, 5, 10, 23}
	sum := 14
	n := len(set)
	if isSubsetSum(set, n, sum) {
		fmt.Println("Found a subset with given sum")
	} else {
		fmt.Println("No subset with given sum")
	}

	subArraySum(set, n, sum)

	arr := []int{10, 2, -2, -20, 10}
	sum = -10
	n = len(arr)
	if isSubsetSum(arr, n, sum) {
		fmt.Println("Found a subset with given sum")
	} else {
		fmt.Println("No subset with given sum")
	}

	subArraySum(arr, n, sum)
}
