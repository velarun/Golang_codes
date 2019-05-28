package main

import "fmt"

func Search(arr []int, n, x int) string {

	if arr[n-1] == x {
		return "Found"
	}

	backup := arr[n-1]
	arr[n-1] = x

	fmt.Println(arr)

	for i := 0; i < n; i++ {

		//this would be executed at-most n times and therefore at-most n comparisons
		if arr[i] == x {

			//replace arr[n-1] with its actual element as in original 'arr[]'
			arr[n-1] = backup

			fmt.Println(arr)

			//if 'x' is found before the '(n-1)th' index, then it is present in the array final comparison
			if i < n-1 {
				return "Found"
			}
		}

	}

	return "Not Found"
}

func main() {

	testlist := []int{4, 6, 1, 5, 8}
	n := len(testlist)
	x := 1
	fmt.Println("Search on Unsorted Array with O(n) complexity:")
	fmt.Println(Search(testlist, n, x))
}
