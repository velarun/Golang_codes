package main

import "fmt"

func printUnion(arr1, arr2 []int) { 
	hs := make(map[int]bool) 

	for _, v := range arr1 { 
		if _, ok := hs[v]; !ok {
			hs[v] = true
		}
	}
	
	for _, v := range arr2 { 
		if _, ok := hs[v]; !ok {
			hs[v] = true
		}
	}

	fmt.Println("Union:") 
    for key := range hs {
        fmt.Print(key, " ")
    }
	fmt.Println() 
}

func printIntersection(arr1, arr2 []int) { 
	hs := make(map[int]bool) 

	for _, v := range arr1 { 
		if _, ok := hs[v]; !ok {
			hs[v] = true
		}
	}
	
	fmt.Println("Intersection:")
	for _, v := range arr2 { 
		if _, ok := hs[v]; ok {
			fmt.Print(v, " ")
		}
	}
	fmt.Println() 
}

func main() { 
	arr1 := []int{7, 1, 5, 2, 3, 6} 
	arr2 := []int{3, 8, 6, 20, 7} 
	printUnion(arr1, arr2) 
	printIntersection(arr1, arr2) 
}