package main

import "fmt"

func BinarySearch(list []int, item int) bool {
	first := 0
	last := len(list) - 1
	found := false

	for first <= last && !found {

		midpoint := (first + last) / 2

		if list[midpoint] == item {
			found = true
		} else {
			if item < list[midpoint] {
				last = last - 1
			} else {
				first = first + 1
			}
		}
	}

	return found

}

func SequentialSearch(list []int, item int) bool {

	found := false

	for i := 0; i < len(list)-1; i++ {
		if list[i] == item {
			found = true
			break
		}
	}

	return found
}

func main() {

	testlist := []int{0, 1, 2, 8, 13, 17, 19, 32, 42}
	item := 13
	fmt.Println("Binary Search:")
	if BinarySearch(testlist, item) {
		fmt.Println(item, "Found")
	} else {
		fmt.Println(item, "Not found")
	}

	item = 4
	if BinarySearch(testlist, item) {
		fmt.Println(item, "Found")
	} else {
		fmt.Println(item, "Not found")
	}

	fmt.Println("Sequential Search:")
	item = 13
	if SequentialSearch(testlist, item) {
		fmt.Println(item, "Found")
	} else {
		fmt.Println(item, "Not found")
	}

	item = 3
	if SequentialSearch(testlist, item) {
		fmt.Println(item, "Found")
	} else {
		fmt.Println(item, "Not found")
	}
}
