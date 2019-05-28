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

func printLeafNodesrighttoleft(root *BinaryNode) {
	// If node is null, return
	if root == nil {
		return
	}

	// If node is leaf node, print its data
	if root.left == nil && root.right == nil {
		fmt.Print(root.data, " ")
		return
	}

	// If right child exists, check for leaf recursively
	if root.right != nil {
		printLeafNodesrighttoleft(root.right)
	}

	// If left child exists, check for leaf recursively
	if root.left != nil {
		printLeafNodesrighttoleft(root.left)
	}
}

func printLeafNodeslefttoright(root *BinaryNode) {
	// If node is null, return
	if root == nil {
		return
	}

	// If node is leaf node, print its data
	if root.left == nil && root.right == nil {
		fmt.Print(root.data, " ")
		return
	}

	// If left child exists, check for leaf recursively
	if root.left != nil {
		printLeafNodeslefttoright(root.left)
	}

	// If right child exists, check for leaf recursively
	if root.right != nil {
		printLeafNodeslefttoright(root.right)
	}
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

	fmt.Println("Printing Leaf nodes from right to left:")
	printLeafNodesrighttoleft(root)
	fmt.Println()
	fmt.Println("Printing Leaf nodes from left to right:")
	printLeafNodeslefttoright(root)
	fmt.Println()
}
