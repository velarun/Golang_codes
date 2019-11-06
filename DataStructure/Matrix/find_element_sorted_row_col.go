package main

import "fmt"

func findElement(arr [][]int, n int, length int) {
    
    i := 0 //col
    j := length - 1 //row
    found := false
    
    for i < length && j >= 0 {
        if arr[i][j] == n {
            fmt.Println("Found")
            found = true
        }
        
        if arr[i][j] > n {
            j--
        } else {
            i++
        }
    }
    
    if !found {
        fmt.Println("Not Found")
    }   
}

func main() {
    arr := [][]int{{10, 20, 30, 40}, {15, 25, 35, 45}, {27, 29, 37, 48}, {32, 33, 39, 50}}
    length := 4

	findElement(arr, 26, length)
	findElement(arr, 27, length)
}