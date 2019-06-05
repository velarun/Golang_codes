package main

import (
	"fmt"
)

//BinaryNode struct
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

var res []int

func SumChildStoreParent(root *BinaryNode) int {

	if root == nil {
		return 0
	}

	local := root.data

	root.data = SumChildStoreParent(root.left) + SumChildStoreParent(root.right)

	return root.data + local
}

func inorderTraversal(node *BinaryNode, res []int) []int {
	if node != nil {
		res = inorderTraversal(node.left, res)
		res = append(res, node.data)
		res = inorderTraversal(node.right, res)
	}
	return res
}

func main() {
	root := &BinaryNode{data: 10, left: nil, right: nil}
	root.left = &BinaryNode{data: -2, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 8, left: nil, right: nil}
	root.left.right = &BinaryNode{data: -4, left: nil, right: nil}
	
	root.right = &BinaryNode{data: 6, left: nil, right: nil}
	root.right.left = &BinaryNode{data: 7, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 5, left: nil, right: nil}
	
	SumChildStoreParent(root)
	fmt.Println("Inorder Traversal: ", inorderTraversal(root, res))
}
