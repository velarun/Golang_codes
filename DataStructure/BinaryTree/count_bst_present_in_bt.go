package main

import (
	"fmt"
	"math"
)

//BinaryNode struct
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

const INT_MIN = math.MinInt64
const INT_MAX = math.MaxInt64
		
func NumberOfBST(root *BinaryNode) (int, int, int, bool) {
	if root == nil { 
		return 0, INT_MIN, INT_MAX, true
	}

	if root.left == nil && root.right == nil {
		return 1, root.data, root.data, true
	}

	L0, L1, L2, L3 := NumberOfBST(root.left) 
	R0, R1, R2, R3 := NumberOfBST(root.right) 

	var bst0, bst1, bst2 int 
	var bst3 bool

	if L3 && R3 && root.data > L1 && root.data < R2 {  
		bst2 = int(math.Min(float64(root.data), math.Min(float64(L2), float64(R2))))
		bst1 = int(math.Max(float64(root.data), math.Max(float64(L1), float64(R1))))
		bst0 = 2 + L0 + R0 

		return bst0, bst1, bst2, bst3 
	}

	bst3 = false
	bst0 = 1 + L0 + R0 

	return bst0, bst1, bst2, bst3 
}

func main() {
	root := &BinaryNode{data: 5, left: nil, right: nil}
	root.left = &BinaryNode{data: 9, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 6, left: nil, right: nil}
	root.left.left.left = &BinaryNode{data: 8, left: nil, right: nil}
	root.left.left.right = &BinaryNode{data: 7, left: nil, right: nil}
	
	root.right = &BinaryNode{data: 3, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 4, left: nil, right: nil}
	
	count, _, _, _ := NumberOfBST(root)
	fmt.Println("Printing number of BST in BT:", count)
}