package main

import "fmt"

//O(n)
func findSplitPoint(arr []int, n int) int { 
	
	leftSum := 0
	for i:=0; i<n; i++ { 
		leftSum += arr[i] 
	}

	rightSum := 0
	for i:=n-1; i>-1; i-- { 
		rightSum += arr[i] 
		leftSum -= arr[i] 
		if rightSum == leftSum { 
			return i 
		}
	}
 
	return -1
}

func printTwoParts(arr []int, n int) { 
	splitPoint := findSplitPoint(arr, n) 

	if splitPoint == -1 || splitPoint == n { 
		fmt.Println("Not Possible") 
		return
	}

	for i:=0; i<n; i++ { 
		if splitPoint == i { 
			fmt.Println() 
		}
		fmt.Print(arr[i], " ")
	}		 
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 5}
	n := len(arr) 
	printTwoParts(arr, n) 
}