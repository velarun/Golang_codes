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

func KLargest(root *BinaryNode, k int) int {

	count := 0

	klarge := 0
	curr := root

	for curr != nil {
		if curr.right == nil {
			count += 1
			if count == k {
				klarge = curr.data
				break
			}
			curr = curr.left
		} else {
			pre := curr.right
			for pre.left != nil && pre.left != curr {
				pre = pre.left
			}

			if pre.left == nil {
				pre.left = curr
				curr = curr.right
			} else {
				pre.left = nil
				count += 1
				if count == k {
					klarge = curr.data
					break
				}
				curr = curr.left
			}
		}
	}
	return klarge
}

func main() {
	root := &BinaryNode{data: 4, left: nil, right: nil}
	root.left = &BinaryNode{data: 2, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 1, left: nil, right: nil}
	root.left.right = &BinaryNode{data: 3, left: nil, right: nil}

	root.right = &BinaryNode{data: 7, left: nil, right: nil}
	root.right.left = &BinaryNode{data: 6, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 10, left: nil, right: nil}

	fmt.Println("kth largest of tree is", KLargest(root, 2))
}
