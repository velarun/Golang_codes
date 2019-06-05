package main

import "fmt"

func getSum(arr []int, n int) int { 

	sum := 0 
	for i:=0; i<n; i++ {
		sum += arr[i] 
	} 	
	return sum
}

//Function to find two missing numbers in range [1, n]. This 
//function assumes that size of array is n-2 and all array 
//elements are distinct 
func findTwoMissingNumbers(arr []int, n int) {

	sum := ((n * (n + 1)) / 2 - getSum(arr, n - 2)) 

	avg := sum/2 

	//Find sum of elements smaller than average (avg) and sum 
	//of elements greater than average (avg) 
	sumSmallerHalf := 0
	sumGreaterHalf := 0

	for i := 0; i < n - 2; i++ { 
		if arr[i] <= avg { 
			sumSmallerHalf += arr[i] 
		} else { 
			sumGreaterHalf += arr[i] 
		}
	}

	fmt.Println("Two Missing Numbers are:") 
	totalSmallerHalf := (avg * (avg + 1)) / 2
	fmt.Println(totalSmallerHalf - sumSmallerHalf) 
	fmt.Println((n * (n + 1)) / 2 - totalSmallerHalf - sumGreaterHalf)

}

func findOneMissingNumber(arr []int, n int) {
	missing := ((n * (n + 1)) / 2 - getSum(arr, n - 1)) 
	fmt.Println("One Missing Number is:")
	fmt.Println(missing)
}

func findOneMissingNoXOR(a []int, n int) { 
	x1 := a[0] 
	x2 := 1
	
	for i:=1; i<n; i++ { 
		x1 = x1 ^ a[i] 
	}

	for i:=2; i<n+2; i++ { 
		x2 = x2 ^ i 
	}

	fmt.Println("One Missing Number using XOR is:")
	fmt.Println(x1 ^ x2)
}

func findTwoMissingNoVisited(arr []int, n int) { 
    mark := make(map[int]bool, n) 
    for i:=0; i<n-2; i++ {
		mark[arr[i]] = true
	}
 
    fmt.Println("Two Missing Numbers using XOR are") 
    for i:=1; i<n+1; i++ { 
        if mark[i] == false { 
			fmt.Print(i, " ")
		}
	} 
  
	fmt.Println()
}

func findTwoMissingNoXOR(arr []int, n int) {

	XOR := arr[0] 
	for i:=1; i<n-2; i++ { 
		XOR ^= arr[i] 
	}

	for i:=1; i<n+1; i++ { 
		XOR ^= i 
	}

	set_bit_no := XOR & -(XOR-1)

	x := 0
	y := 0

	for i:=0; i<n-2; i++ {
		if (arr[i] & set_bit_no) != 0 { 
			x = x ^ arr[i] 
		} else { 
			y = y ^ arr[i] 
		}
	}

	for i:=1; i<n+1; i++ { 
		if (i & set_bit_no) != 0 { 
			x = x ^ i 				
		} else {
			y = y ^ i 
		}
	}

	fmt.Println("Two Missing Numbers are:", x, y) 
}

func main() {
	arr := []int{1, 3, 5, 6}

	//two missing number
	n := 2 + len(arr) 
	findTwoMissingNumbers(arr, n) 

	findTwoMissingNoVisited(arr, n)

	findTwoMissingNoXOR(arr, n)

	arr = []int{1, 2, 3, 5, 6}
	n = 1 + len(arr) 

	//one missing number
	findOneMissingNumber(arr, n) 

	//One missing number using xor
	findOneMissingNoXOR(arr, len(arr))

}
