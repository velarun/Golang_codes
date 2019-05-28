package main

import (
	"fmt"
	"sort"
)

//BinaryNode struct
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

func getVerticalOrder(root *BinaryNode, hd int, m map[int][]*BinaryNode) {

	if root == nil {
		return
	}

	obj := &BinaryNode{data: root.data, left: root.left, right: root.right}
	m[hd] = append(m[hd], obj)

	//Store nodes in left subtree
	getVerticalOrder(root.left, hd-1, m)

	//Store nodes in right subtree
	getVerticalOrder(root.right, hd+1, m)
}

func printVerticalOrder(root *BinaryNode) {

	m := make(map[int][]*BinaryNode)
	var hd int
	getVerticalOrder(root, hd, m)

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, v1 := range keys {
		for _, v2 := range m[v1] {
			fmt.Print(v2.data, " ")
		}
		fmt.Println()
	}
}

func main() {
	root := &BinaryNode{data: 1, left: nil, right: nil}
	root.left = &BinaryNode{data: 2, left: nil, right: nil}
	root.right = &BinaryNode{data: 3, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 4, left: nil, right: nil}
	root.left.right = &BinaryNode{data: 5, left: nil, right: nil}
	root.right.left = &BinaryNode{data: 6, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 7, left: nil, right: nil}
	root.right.left.right = &BinaryNode{data: 8, left: nil, right: nil}
	root.right.right.right = &BinaryNode{data: 9, left: nil, right: nil}

	fmt.Println("Print Vertical Order of given binary tree is")
	printVerticalOrder(root)
}
