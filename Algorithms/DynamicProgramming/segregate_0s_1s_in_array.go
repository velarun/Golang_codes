package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1,0,1,0,1,1}
	leng := len(arr1)
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
