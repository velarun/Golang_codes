package main

import "fmt"

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

func InsertLevelOrder(arr []int, root *BinaryNode, i, n int) *BinaryNode {

	if i < n {
		temp := &BinaryNode{data: arr[i], left: nil, right: nil}
		root = temp
		root.left = InsertLevelOrder(arr, root.left, 2*i+1, n)
		root.right = InsertLevelOrder(arr, root.right, 2*i+2, n)
	}

	return root
}

func main() {
	n := 10

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}

	root := &BinaryNode{}
	root = InsertLevelOrder(arr, root, 0, n)

	fmt.Println("Construct binary tree of size 10 & Printing root.left.left.right:")
	fmt.Println(root.left.left.right)
}
