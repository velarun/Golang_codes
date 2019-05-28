package main

import (
	"fmt"
)

//BinaryNode struct
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

func fullBinaryTree(root *BinaryNode) bool {

	if root == nil {
		return true
	}

	if root.left == nil && root.right == nil {
		return true
	}

	if root.left != nil && root.right != nil {
		return fullBinaryTree(root.left) && fullBinaryTree(root.right)
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

	if fullBinaryTree(root) {
		fmt.Println("Tree is Full Binary Tree")
	} else {
		fmt.Println("Tree is not a Full Binary Tree")
	}

	root = &BinaryNode{data: 8, left: nil, right: nil}
	root.left = &BinaryNode{data: 15, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 1, left: nil, right: nil}
	root.left.right = &BinaryNode{data: 6, left: nil, right: nil}
	root.left.right.left = &BinaryNode{data: 4, left: nil, right: nil}
	root.left.right.right = &BinaryNode{data: 7, left: nil, right: nil}

	root.right = &BinaryNode{data: 10, left: nil, right: nil}
	root.right.left = &BinaryNode{data: 14, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 13, left: nil, right: nil}

	if fullBinaryTree(root) {
		fmt.Println("Tree is Full Binary Tree")
	} else {
		fmt.Println("Tree is not a Full Binary Tree")
	}
}
