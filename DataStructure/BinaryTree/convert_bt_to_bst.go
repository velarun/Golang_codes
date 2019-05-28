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

func countNodes(root *BinaryNode) int {
	if root == nil {
		return 0
	}

	return countNodes(root.left) + countNodes(root.right) + 1
}

func arrayToBST(arr *[]int, root *BinaryNode) {

	if root == nil {
		return
	}

	arrayToBST(arr, root.left)

	root.data = (*arr)[0]
	(*arr) = (*arr)[1:]

	arrayToBST(arr, root.right)
}

func binaryTreeToBST(root *BinaryNode) {

	if root == nil {
		return
	}

	//Count the number of nodes in Binary Tree so that
	//we know the size of temporary array to be created
	n := countNodes(root)
	fmt.Println("Number of nodes: ", n)

	//Create the temp array and store the inorder traveral of tree
	arr := []int{}
	arr = inorderTraversal(root, arr)

	//Sort the array
	sort.Ints(arr)

	//copy array elements back to binary tree
	arrayToBST(&arr, root)
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
	root.left = &BinaryNode{data: 30, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 20, left: nil, right: nil}

	root.right = &BinaryNode{data: 15, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 5, left: nil, right: nil}

	binaryTreeToBST(root)

	fmt.Println("Inorder traversal of the converted Binary Search Tree:", inorderTraversal(root, res))
}
