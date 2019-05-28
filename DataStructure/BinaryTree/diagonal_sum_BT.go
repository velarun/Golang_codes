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

func diagonalSumUtil(root *BinaryNode, vd int, diagonalSum map[int]int64) *BinaryNode {

	if root == nil {
		return nil
	}

	if _, ok := diagonalSum[vd]; !ok {
		diagonalSum[vd] = 0
	}
	diagonalSum[vd] += root.data

	diagonalSumUtil(root.left, vd+1, diagonalSum)

	diagonalSumUtil(root.right, vd, diagonalSum)

	return root
}

func diagonalSum(root *BinaryNode) {

	diagonalSum := make(map[int]int64)

	diagonalSumUtil(root, 0, diagonalSum)

	fmt.Println("Diagonal sum in a binary tree is:")

	for i := 0; i < len(diagonalSum); i++ {
		fmt.Println(diagonalSum[i])
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

	diagonalSum(root)
}
