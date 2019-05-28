package main

import "fmt"

func main() {
	arr := []int{6, 2, 4, 18, 3, 8, 9, 10, 11}

	set := make(map[int]bool)
	for _, v := range arr {
		set[v] = true
	}

	max := 0
	looper := 0
	for _, v := range arr {
		looper++
		v++
		for {
			if !set[v] {
				break
			}
			v++
			looper++
		}

		if looper > max {
			max = looper
		}
		looper = 0
	}

	fmt.Println(max)

}
