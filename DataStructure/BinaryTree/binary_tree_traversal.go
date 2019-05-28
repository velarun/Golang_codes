package main

import (
	"fmt"
	"io"
	"os"
)

//BinaryNode struct
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

var res []int64

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func inorderTraversal(node *BinaryNode, res []int64) []int64 {
	if node != nil {
		res = inorderTraversal(node.left, res)
		res = append(res, node.data)
		res = inorderTraversal(node.right, res)
	}
	return res
}

func preorderTraversal(node *BinaryNode, res []int64) []int64 {
	if node != nil {
		res = append(res, node.data)
		res = preorderTraversal(node.left, res)
		res = preorderTraversal(node.right, res)
	}
	return res
}

func postorderTraversal(node *BinaryNode, res []int64) []int64 {
	if node != nil {
		res = postorderTraversal(node.left, res)
		res = postorderTraversal(node.right, res)
		res = append(res, node.data)
	}
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

	fmt.Println("Tree:")
	print(os.Stdout, root, 0, 'M')
	fmt.Println("Inorder Traversal: ", inorderTraversal(root, res))
	fmt.Println("Preorder Traversal: ", preorderTraversal(root, res))
	fmt.Println("Postorder Traversal: ", postorderTraversal(root, res))
}
