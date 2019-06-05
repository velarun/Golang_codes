package main

import "fmt"

func printRotate180Clockwise(A [][]int) {
	N := len(A[0])
	for i:=N-1; i>=0; i-- {
		for j:=N-1; j>=0; j-- {
			fmt.Print(A[i][j], " ")
		}
		fmt.Println()
	}
}

func rotate90Clockwise(A [][]int) { 
	N := len(A[0]) 
	for i:=0; i<N/2; i++ {
		for j:=i; j<N-i-1; j++ { 
			temp := A[i][j] 
			A[i][j] = A[N - 1 - j][i] 
			A[N - 1 - j][i] = A[N - 1 - i][N - 1 - j] 
			A[N - 1 - i][N - 1 - j] = A[j][N - 1 - i] 
			A[j][N - 1 - i] = temp 
		}
	}
}

func printMatrix(A [][]int) { 
    N := len(A[0]) 
    for i:=0; i<N; i++ { 
		fmt.Println(A[i])
	}
}

func main() { 
	A := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	printRotate180Clockwise(A)
	A = [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	rotate90Clockwise(A)
	printMatrix(A) 
}