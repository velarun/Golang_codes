//O(n^2)
package main

import (
	"fmt"
	"math"
)

//BinaryNode struct
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

func Height(node *BinaryNode) int {

	if node == nil {
		return 0
	}

	//If tree is not empty then height = 1 + max of left
	//height and right heights
	return 1 + int(math.Max(float64(Height(node.left)), float64(Height(node.right))))
}

func Diameter(root *BinaryNode) int {

	if root == nil {
		return 0
	}

	//Get the height of left and right sub-trees
	lheight := Height(root.left)
	//fmt.Println("lh: ", lheight)
	rheight := Height(root.right)
	//fmt.Println("rh: ", rheight)

	//Get the diameter of left and irgh sub-trees
	ldiameter := Diameter(root.left)
	//fmt.Println("ld: ", ldiameter)
	rdiameter := Diameter(root.right)
	//fmt.Println("rd: ", rdiameter)

	//Return max of the following tree:
	//1) Diameter of left subtree
	//2) Diameter of right subtree
	//3) Height of left subtree + height of right subtree +1
	return int(math.Max(float64(lheight+rheight+1), math.Max(float64(ldiameter), float64(rdiameter))))
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

	fmt.Println("Diameter of given binary tree is", Diameter(root))
}
