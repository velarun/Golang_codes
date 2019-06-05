package main

import (
	"fmt"
)

func segregate0and1count(arr1 []int, leng int) {
	count := 0
	
	for i:=0 ; i<leng ; i++ {
		if arr1[i] == 0 {
			count+=1
		}
	}
	
	for i:=0 ; i<count ; i++ {
		arr1[i] = 0 
	}
	
	for i:=count ; i<leng ; i++ {
		arr1[i] = 1 
	}
	
	fmt.Println(arr1)
}

func segregate0and1twoindex(arr []int, size int) { 
	//Initialize left and right indexes 
	left, right := 0, size-1
	
	for left < right {
		//Increment left index while we see 0 at left 
		for arr[left] == 0 && left < right { 
			left++
		}

		//Decrement right index while we see 1 at right 
		for arr[right] == 1 && left < right { 
			right--
		}

		//If left is smaller than right then there is a 1 at left 
		//and a 0 at right. Exchange arr[left] and arr[right] 
		if left < right {
			arr[left] = 0
			arr[right] = 1
			left++
			right--
		}
	}

	fmt.Println(arr) 
}

func segregate0and1twopointer(arr []int, size int) { 

	type0 := 0
	type1 := size - 1
	
	for type0 < type1 { 
		if arr[type0] == 1 { 
			arr[type0], arr[type1] = arr[type1], arr[type0]
			type1--
		} else {
			type0++
		}
	}

	fmt.Println(arr)
}

func main() {
	arr1 := []int{1,0,1,0,1,1}
	leng := len(arr1)
	
	fmt.Println(arr1)
	segregate0and1count(arr1, leng)
	fmt.Println(arr1) //without pointing pointer array saves values

	arr1 = []int{1,0,1,0,1,1}
	leng = len(arr1)
	fmt.Println(arr1)
	segregate0and1twoindex(arr1, leng)

	arr1 = []int{1,0,1,0,1,1}
	leng = len(arr1)
	fmt.Println(arr1)
	segregate0and1twopointer(arr1, leng)

}