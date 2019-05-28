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

var res []int

type tswaps []tswap

type tswap struct {
	first  int
	second int
}

type bySort []tswap

func (b bySort) Len() int {
	return len(b)
}

func (b bySort) Swap(i, j int) {
	b[i].first, b[j].first = b[j].first, b[i].first
	b[i].second, b[j].second = b[j].second, b[i].second
}

func (b bySort) Less(i, j int) bool {
	return b[i].first < b[j].first
}

// Function to find minimum swaps to sort an array
func minSwaps(v []int) int {

	t := make([]tswap, len(v))

	ans := 0
	for i := 0; i < len(v); i++ {
		t[i].first = v[i]
		t[i].second = i
	}

	sort.Sort(bySort(t))

	for i := 0; i < len(t); i++ {
		if i == t[i].second {
			continue
		} else {
			bySort.Swap(bySort(t), i, t[i].second)
		}

		// Second is not equal to i
		if i != t[i].second {
			i -= 1
		}

		ans++
	}
	return ans
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
	root := &BinaryNode{data: 5, left: nil, right: nil}
	root.left = &BinaryNode{data: 6, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 8, left: nil, right: nil}
	root.left.right = &BinaryNode{data: 9, left: nil, right: nil}

	root.right = &BinaryNode{data: 7, left: nil, right: nil}
	root.right.left = &BinaryNode{data: 10, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 11, left: nil, right: nil}

	inOrder := inorderTraversal(root, res)

	fmt.Println("Minimum Swap to convert Binary Tree to Binary Search Tree:", minSwaps(inOrder))
}
