package main

import "fmt"

func main() {
	arr := []int{28, 33, 16, 7, 5, 88}

	max := arr[0] // assume first value is the smallest

	for _, value := range arr {
		if value > max {
			max = value // found another smaller value, replace previous value in max
		}
	}

	fmt.Println("The biggest/largest value is : ", max)

	largest := arr[0]
	second := arr[0]

	for _, i := range arr {
		if (i > largest) {
			second = largest;
			largest = i;
		} else if (i > second) {
		      	second = i;
		}
	}

	fmt.Println("The largest value is: ", largest, "and the Second largest value is: ", second)
}
