package main

import (
	"fmt"
	"math"
)

const INT_MIN = math.MinInt64
const INT_MAX = math.MaxInt64

type BinaryNodeDetails struct {
	data int
	min int
	max int
}

func levelOrderIsOfBST(arr []int, n int) bool { 
	if n == 0 { 
		return true
	}

	var q []BinaryNodeDetails
	
	//index variable to access array elements 
	i := 0
	
	newNode := BinaryNodeDetails{data: arr[i], min: INT_MIN, max: INT_MAX}
	i += 1
	q = append(q, newNode) 
	 
	for i != n && len(q) != 0 {	
		temp := q[0] 
		q = q[1:]
	
		if i < n && (arr[i] < temp.data && arr[i] > temp.min) {  
			newNode = BinaryNodeDetails{data: arr[i], min: temp.min, max: temp.data} 
			i += 1
			q = append(q, newNode)
		}			 

		if i < n && (arr[i] > temp.data && arr[i] < temp.max) {
			newNode = BinaryNodeDetails{data: arr[i], min: temp.data, max: temp.max} 
			i += 1
			q = append(q, newNode)
		}		 
	}

	if i == n {
		return true
	}

	return false
}		

func main() {
	arr1 := []int{7, 4, 12, 3, 6, 8, 1, 5, 10} 
	n1 := len(arr1)	 
	if levelOrderIsOfBST(arr1, n1) { 
		fmt.Println("Yes") 
	} else { 
		fmt.Println("No") 
	}

	arr2 := []int{11, 6, 13, 5, 12, 10}
	n2 := len(arr2)	 
	if levelOrderIsOfBST(arr2, n2) { 
		fmt.Println("Yes") 
	} else { 
		fmt.Println("No") 
	}
}