package main

import "fmt"

func count(S []int, m int, n int) {

	if n == 0 {
		return
	}

	if n < 0 {
		return
	}

	arr := []int{}

	for i := m - 1; i >= 0; i-- {
		for n >= S[i] {
			n -= S[i]
			arr = append(arr, S[i])
		}
	}

	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {

	var c []int = []int{1, 10, 20, 50, 100}
	var n int = 196
	count(c, len(c), n)
}
