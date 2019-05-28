package main

import "fmt"

func BubbleSort(a []int, n int) {
	for j := n - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if a[i] > a[i+1] {
				temp := a[i]
				a[i] = a[i+1]
				a[i+1] = temp
			}
			fmt.Println(a)
		}
	}
}

func InsertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
		fmt.Println(arr)
	}
}

func SelectionSort(alist []int, n int) {

	for fillslot := n - 1; fillslot > 0; fillslot-- {
		positionOfMax := 0
		for location := 1; location < fillslot+1; location++ {
			if alist[location] > alist[positionOfMax] {
				positionOfMax = location
			}
		}
		fmt.Println(alist)
		temp := alist[fillslot]
		alist[fillslot] = alist[positionOfMax]
		alist[positionOfMax] = temp
	}
}

func QuickSort(alist []int) {
	quickSortHelper(alist, 0, len(alist)-1)
}

func quickSortHelper(alist []int, first, last int) {
	fmt.Println(alist)
	if first < last {
		splitpoint := partition(alist, first, last)

		quickSortHelper(alist, first, splitpoint-1)
		quickSortHelper(alist, splitpoint+1, last)
	}
}

func partition(alist []int, first, last int) int {
	pivotvalue := alist[first]

	leftmark := first + 1
	rightmark := last

	var done bool
	for !done {
		for leftmark <= rightmark && alist[leftmark] <= pivotvalue {
			leftmark = leftmark + 1
		}

		for alist[rightmark] >= pivotvalue && rightmark >= leftmark {
			rightmark = rightmark - 1
		}

		if rightmark < leftmark {
			done = true
		} else {
			temp := alist[leftmark]
			alist[leftmark] = alist[rightmark]
			alist[rightmark] = temp
		}
	}

	temp := alist[first]
	alist[first] = alist[rightmark]
	alist[rightmark] = temp

	return rightmark
}

func main() {
	alist := []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
	fmt.Println("Unsorted Array:")
	fmt.Println(alist)
	fmt.Println("=================")
	BubbleSort(alist, len(alist))
	fmt.Println("=================")
	fmt.Println("Sorted Array using bubble sort:")
	fmt.Println(alist)

	alist = []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
	fmt.Println("Unsorted Array:")
	fmt.Println(alist)
	fmt.Println("=================")
	InsertionSort(alist, len(alist))
	fmt.Println("=================")
	fmt.Println("Sorted Array using insertion sort:")
	fmt.Println(alist)

	alist = []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
	fmt.Println("Unsorted Array:")
	fmt.Println(alist)
	fmt.Println("=================")
	SelectionSort(alist, len(alist))
	fmt.Println("=================")
	fmt.Println("Sorted Array using selection sort:")
	fmt.Println(alist)

	alist = []int{54, 26, 93, 17, 77, 31, 44, 55, 20}
	fmt.Println("Unsorted Array:")
	fmt.Println(alist)
	fmt.Println("=================")
	QuickSort(alist)
	fmt.Println("=================")
	fmt.Println("Sorted Array using Quick sort:")
	fmt.Println(alist)

}
