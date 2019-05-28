package main

import (
	"fmt"
	"math"
)

//BinaryNode struct
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

func height(root *BinaryNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(float64(height(root.left)), float64(height(root.right))) + 1)
}

func isBalanced(root *BinaryNode) bool {

	if root == nil {
		return true
	}

	//for left and right subtree height
	lh := height(root.left)
	rh := height(root.right)

	//allowed values for (lh - rh) are 1, -1, 0
	if (int(math.Abs(float64(lh-rh))) <= 1) && isBalanced(root.left) == true && isBalanced(root.right) == true {
		return true
	}

	return false
}

func main() {
	root := &BinaryNode{data: 8, left: nil, right: nil}
	root.left = &BinaryNode{data: 15, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 1, left: nil, right: nil}
	root.left.right = &BinaryNode{data: 6, left: nil, right: nil}
	root.left.right.left = &BinaryNode{data: 4, left: nil, right: nil}
	root.left.right.right = &BinaryNode{data: 7, left: nil, right: nil}

	root.right = &BinaryNode{data: 10, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 14, left: nil, right: nil}
	root.right.right.left = &BinaryNode{data: 13, left: nil, right: nil}

	if isBalanced(root) {
		fmt.Println("Tree is balanced")
	} else {
		fmt.Println("Tree is not balanced")
	}

	root = &BinaryNode{data: 8, left: nil, right: nil}
	root.left = &BinaryNode{data: 15, left: nil, right: nil}
	root.right = &BinaryNode{data: 10, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 1, left: nil, right: nil}
	root.left.right = &BinaryNode{data: 6, left: nil, right: nil}
	root.right.left = &BinaryNode{data: 10, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 14, left: nil, right: nil}

	if isBalanced(root) {
		fmt.Println("Tree is balanced")
	} else {
		fmt.Println("Tree is not balanced")
	}
}
