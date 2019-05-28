//Find n’th number in a number system with only 3 and 4
package main

import "fmt"

func Find(n int) {

	arr := make([]string, n+1)

	size := 1
	m := 1

	for size <= n {
		for i := 0; i < m && (size+i) <= n; i++ {
			arr[size+i] = "3" + arr[size-m+i]
		}

		for i := 0; i < m && (size+m+i) <= n; i++ {
			arr[size+m+i] = "4" + arr[size-m+i]
		}

		m = m << 1
		//m = m*2
		size = size + m
	}

	fmt.Println(arr[n])
}

func main() {
	for i := 1; i < 16; i++ {
		Find(i)
	}
}
