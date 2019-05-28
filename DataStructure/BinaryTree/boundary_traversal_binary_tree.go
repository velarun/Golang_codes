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

func printLeaves(root *BinaryNode) {
	if root != nil {
		printLeaves(root.left)
		if root.left == nil && root.right == nil {
			fmt.Print(root.data, " ")
		}
		printLeaves(root.right)
	}
}

//A function to print all left boundary nodes, except a
//leaf node. Print the nodes in TOP DOWN manner
func printBoundaryLeft(root *BinaryNode) {
	if root != nil {
		if root.left != nil {
			//to ensure top down order, print the node
			//before calling itself for left subtree
			fmt.Print(root.data, " ")
			printBoundaryLeft(root.left)
		} else if root.right != nil {
			fmt.Print(root.data, " ")
			printBoundaryLeft(root.right)
		}
		//do nothing if it is a leaf node, this way we
		//avoid duplicates in output
	}
}

//A function to print all right boundary nodes, except
//a leaf node. Print the nodes in BOTTOM UP manner
func printBoundaryRight(root *BinaryNode) {
	if root != nil {
		if root.right != nil {
			//to ensure bottom up order, first call for
			//right subtree, then print this node
			printBoundaryRight(root.right)
			fmt.Print(root.data, " ")
		} else if root.left != nil {
			printBoundaryRight(root.left)
			fmt.Print(root.data, " ")
		}
		//do nothing if it is a leaf node, this way we
		//avoid duplicates in output
	}
}

func Boundary(root *BinaryNode) {
	if root != nil {
		fmt.Print(root.data, " ")

		//Print the left boundary in top-down manner
		printBoundaryLeft(root.left)

		//Print all leaf nodes
		printLeaves(root.left)
		printLeaves(root.right)

		//Print the right boundary in bottom-up manner
		printBoundaryRight(root.right)
	}
}

func main() {
	root := &BinaryNode{data: 20, left: nil, right: nil}
	root.left = &BinaryNode{data: 8, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 4, left: nil, right: nil}
	root.left.right = &BinaryNode{data: 12, left: nil, right: nil}
	root.left.right.left = &BinaryNode{data: 10, left: nil, right: nil}
	root.left.right.right = &BinaryNode{data: 14, left: nil, right: nil}

	root.right = &BinaryNode{data: 22, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 25, left: nil, right: nil}

	fmt.Println("Boundary Traversal of binary tree:")
	Boundary(root)
	fmt.Println()
}
