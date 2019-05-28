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

func printLevelOrder(root *BinaryNode) {

	if root == nil {
		return
	}

	q := []*BinaryNode{}

	//Enqueue Root and Initialize Height
	q = append(q, root)

	for len(q) > 0 {
		node := q[0]
		fmt.Print(node.data, " ")
		q = q[1:]

		if node.left != nil {
			q = append(q, node.left)
		}

		if node.right != nil {
			q = append(q, node.right)
		}
	}
	fmt.Println()
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

	fmt.Println("Level order traversal of binary tree is -")
	printLevelOrder(root)
}
