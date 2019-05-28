package main

import "fmt"

func Prime(n int) bool {
	flag := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			flag = false
			break
		}
	}

	return flag
}

func Closest(n int) {

	incrs := -1
	multi := -1
	count := 1

	if n < 2 {
		fmt.Println(n)
	}

	for {
		if Prime(n) {
			fmt.Println(n)
			break
		} else {
			n += incrs
			multi *= -1
			count += 1
			incrs = multi * count
		}
	}

	return
}

func main() {
	n := 11
	Closest(n)
}
