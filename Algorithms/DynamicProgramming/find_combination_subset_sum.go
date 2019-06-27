package main

import "fmt"

func sum(arr []int) int {
		sum := 0
		for _, v := range arr {
				sum += v
		}

		return sum
}

func subset_sum(numbers []int, target int, partial []int) {
	
		s := sum(partial)

    //check if the partial sum is equals to target
    if s == target { 
				fmt.Println(partial, target)
		}

		if s >= target {
				return 
		}

    for i:=0; i<len(numbers); i++ {
				n := numbers[i]
				remaining := numbers[i+1:]
				partial_rec := append(partial, n)
				subset_sum(remaining, target, partial_rec) 
		}
}

func subset(numbers []int, target int) {
		var partial []int
		subset_sum(numbers, target, partial)
}

func main() {
		set := []int{3,9,8,4,5,7,10}
		sum := 10
		subset(set, sum)
}