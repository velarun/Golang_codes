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

func findLCA(root *BinaryNode, n1, n2 int64) *BinaryNode {

	if root == nil {
		return nil
	}

	if root.data == n1 || root.data == n2 {
		return root
	}

	left_lca := findLCA(root.left, n1, n2)
	right_lca := findLCA(root.right, n1, n2)

	if left_lca != nil && right_lca != nil {
		return root
	}

	if left_lca != nil {
		return left_lca
	} else {
		return right_lca
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

	n1 := 4
	n2 := 10
	res := findLCA(root, 4, 10).data
	fmt.Printf("Lowest common ancestor between %d and %d is: %d\n", n1, n2, res)
}
