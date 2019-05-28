package main

import "fmt"

func RemoveDuplicates(arr []int, n int) int {

	if n == 0 || n == 1 {
		return n
	}

	j := 0

	for i := 0; i < n-1; i++ {
		fmt.Println(j, arr, arr[i], arr[i+1])
		if arr[i] != arr[i+1] {
			arr[j] = arr[i]
			j++
		}
	}

	arr[j] = arr[n-1]
	j++

	return j
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 4, 5, 5}
	n := len(arr)

	n = RemoveDuplicates(arr, n)
	fmt.Println(arr)
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
