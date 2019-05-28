package main

import "fmt"

func product(arr []int) int {
	var prod = 1
	for _, v := range arr {
		prod *= v
	}

	return prod
}

func ProductArray(a []int) []int {

	b := make([]int, len(a))
	b[0] = product(a[1:])

	for n := 1; n < len(a); n++ {
		b[n] = b[n-1] * a[n-1] / a[n]
	}

	return b
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(ProductArray(a))
}
