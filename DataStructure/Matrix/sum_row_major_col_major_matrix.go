package main

import "fmt"

func SN(n, an int) int {
	return (n * (1 + an)) / 2
}

//O(n)
func Trace(n, m int) int {

	an := 1 + (n-1)*(m+1)
	rowmajorSum := SN(n, an)
	an = 1 + (n-1)*(n+1)
	colmajorSum := SN(n, an)

	return rowmajorSum + colmajorSum
}

//O(n^2)
func TraceComplex(n, m int) int {

	A := make([][]int, m)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
	}

	B := make([][]int, m)
	for i := 0; i < n; i++ {
		B[i] = make([]int, m)
	}

	C := make([][]int, m)
	for i := 0; i < n; i++ {
		C[i] = make([]int, m)
	}

	cnt := 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			A[i][j] = cnt
			cnt++
		}
	}

	cnt = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			B[j][i] = cnt
			cnt++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}

	var sum int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == j {
				sum += C[i][j]
			}
		}
	}

	return sum
}

func main() {
	N := 3
	M := 3
	fmt.Println(Trace(N, M))
	fmt.Println(TraceComplex(N, M))
}
