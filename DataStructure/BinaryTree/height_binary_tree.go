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

func TreeHeight(root *BinaryNode) int {

	if root == nil {
		return 0
	}

	q := []*BinaryNode{}

	//Enqueue Root and Initialize Height
	q = append(q, root)
	height := 0

	for {
		nodeCount := len(q)

		if nodeCount == 0 {
			break
		}

		height += 1

		//Dequeue all nodes of current level and Enqueue all nodes of next level
		for nodeCount > 0 {
			node := q[0]
			q = q[1:]

			if node.left != nil {
				q = append(q, node.left)
			}

			if node.right != nil {
				q = append(q, node.right)
			}

			nodeCount -= 1
		}
	}

	return height
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

	fmt.Println("Height of tree is", TreeHeight(root))
}
