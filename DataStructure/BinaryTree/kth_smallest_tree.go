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

	ksmall := 0
	curr := root

	for curr != nil {
		if curr.left == nil {
			count += 1
			if count == k {
				ksmall = curr.data
				break
			}
			curr = curr.right
		} else {
			pre := curr.left
			for pre.right != nil && pre.right != curr {
				pre = pre.right
			}

			if pre.right == nil {
				pre.right = curr
				curr = curr.left
			} else {
				pre.right = nil
				count += 1
				if count == k {
					ksmall = curr.data
					break
				}
				curr = curr.right
			}
		}
	}
	return ksmall
}

func main() {
	root := &BinaryNode{data: 50, left: nil, right: nil}
	root.left = &BinaryNode{data: 30, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 20, left: nil, right: nil}
	root.left.right = &BinaryNode{data: 40, left: nil, right: nil}

	root.right = &BinaryNode{data: 70, left: nil, right: nil}
	root.right.left = &BinaryNode{data: 60, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 80, left: nil, right: nil}

	fmt.Println("kth smallest of tree is", KLargest(root, 2))
}
