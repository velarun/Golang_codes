package main

import (
	"fmt"
	"io"
)

//BinaryNode struct
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

const _MIN = -2147483648
const _MAX = 2147483648

func max(a, b int64) int64 {
	if a <= b {
		return b
	}

	return a
}

func min(a, b int64) int64 {
	if a >= b {
		return b
	}

	return a
}

func maxDiffUtil(node *BinaryNode, res int64) (int64, int64) {

	if node == nil {
		return _MAX, res
	}

	if node.left == nil && node.right == nil {
		return node.data, res
	}

	a, res := maxDiffUtil(node.left, res)
	b, res := maxDiffUtil(node.right, res)
	val := min(a, b)

	res = max(res, node.data-val)

	return min(val, node.data), res
}

func maxDiff(node *BinaryNode) int64 {

	var res int64 = _MIN
	_, res = maxDiffUtil(node, res)
	return res
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

	fmt.Printf("Maximum difference between a node and its ancestor is : %d\n", maxDiff(root))

}
