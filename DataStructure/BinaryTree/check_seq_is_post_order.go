package main

import "fmt"

func solve(arr []int, strt int, end int) bool {
	
	if strt>=end {
		return true
	}

	x:=arr[end]
	index:=end-1
	for index>=strt && arr[index]>x {
		index--
	}

	temp:=index
	for index>=strt {
		if arr[index]>x {
			return false
		}
		index--
	}

	if temp==strt {
		return solve(arr,strt,end-1)
	} else{
		return (solve(arr,strt,temp) && solve(arr,temp+1,end-1))
	}
}

func main() {
	arr1 := []int{6, 9, 7, 14, 31, 12, 10}
	if solve(arr1, 0, len(arr1)-1) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	arr2 := []int{1, 2, 5, 6, 7, 3, 8, 9, 10, 4}
	if solve(arr2, 0, len(arr2)-1) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
