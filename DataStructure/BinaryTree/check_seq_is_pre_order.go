package main 

import (
	"fmt"
	"math"
)

const MIN_VALUE = math.MinInt64
const MAX_VALUE = math.MaxInt64

func main() {

	preorder1 := []int{6,3,0,5,4,8,7,9}
	if isBstPossible(preorder1, MIN_VALUE, MAX_VALUE, 0, len(preorder1)-1) { 
		fmt.Println("Yes") 
	} else { 
		fmt.Println("No") 
	}

	preorder2 := []int{1,2,0,3,5}
	if isBstPossible(preorder2, MIN_VALUE, MAX_VALUE, 0, len(preorder2)-1) { 
		fmt.Println("Yes") 
	} else { 
		fmt.Println("No") 
	}
}

func isBstPossible(preorder []int, min int, max int, li int, ri int) bool {
	if li==ri && preorder[li]>min && preorder[li]<max { 
		return true
	}

	if preorder[li]<=min || preorder[li]>=max { 
		return false
	}

	var lsi int 
	if preorder[li+1] < preorder[li] {
		lsi = li+1
	} else {
		lsi = -1
	}
	rsi := findNextHigherElementIndex(preorder, li, ri)

	var rsii int
	if rsi-1 > lsi {
		rsii = rsi-1
	} else {
		rsii = lsi
	}

	isLeftValid := (lsi==-1 || isBstPossible(preorder, min, preorder[li], lsi, rsii))
	isRightValid := (rsi==-1 || isBstPossible(preorder, preorder[li], max, rsi, ri))
	
	return isLeftValid && isRightValid
}

func findNextHigherElementIndex(preorder []int, li, ri int) int {
	x := -1
	for i := li+1; i <= ri ; i++ {
		if preorder[i] > preorder[li] {
			x = i
			break
		}
	}
	return x
}