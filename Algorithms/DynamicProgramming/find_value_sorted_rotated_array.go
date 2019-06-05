package main

import "fmt"

func search(arr []int, l, h, key int) int {
	if l > h { 
		return -1
	}
	
	mid := (l + h) / 2
	if arr[mid] == key { 
		return mid 
	}

	if arr[l] <= arr[mid] { 
		//As this subarray is sorted, we can quickly 
		//check if key lies in half or other half 
		if key >= arr[l] && key <= arr[mid] { 
			return search(arr, l, mid-1, key) 
		}
		return search(arr, mid+1, h, key) 
	}

	//If arr[l..mid] is not sorted, then arr[mid... r] 
	//must be sorted 
	if key >= arr[mid] && key <= arr[h] { 
		return search(arr, mid+1, h, key) 
	}

	return search(arr, l, mid-1, key) 
}

def findMin(arr, low, high): 
    # This condition is needed to handle the case when array is not 
    # rotated at all 
    if high < low: 
        return arr[0] 
  
    # If there is only one element left 
    if high == low: 
        return arr[low] 
  
    # Find mid 
    mid = int((low + high)/2) 
  
    # Check if element (mid+1) is minimum element. Consider 
    # the cases like [3, 4, 5, 1, 2] 
    if mid < high and arr[mid+1] < arr[mid]: 
        return arr[mid+1] 
  
    # Check if mid itself is minimum element 
    if mid > low and arr[mid] < arr[mid - 1]: 
        return arr[mid] 
  
    # Decide whether we need to go to left half or right half 
    if arr[high] > arr[mid]: 
        return findMin(arr, low, mid-1) 
    return findMin(arr, mid+1, high) 

func main() { 
	arr := []int{4, 5, 6, 7, 8, 9, 1, 2, 3} 
	key := 6
	i := search(arr, 0, len(arr)-1, key) 
	if i != -1 { 
		fmt.Printf("Index: %d\n", i) 
	} else { 
		fmt.Println("Key not found") 
	}
}