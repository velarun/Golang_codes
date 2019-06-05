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

func printSpiral(root *BinaryNode) {
	if root == nil { 
		return 
	} 

	//For levels to be printed from right to left
	var s1 []*BinaryNode
	//For levels to be printed from left to right 
	var s2 []*BinaryNode
 
	s1 = append(s1, root) 
 
	for len(s1) != 0 || len(s2) != 0 { 
		for len(s1) != 0 { 
			temp := s1[len(s1)-1] 
			s1 = s1[:len(s1)-1] 
			fmt.Print(temp.data, " ") 

			if temp.right != nil { 
				s2 = append(s2, temp.right) 
			}

			if temp.left != nil {
				s2 = append(s2, temp.left)
			}
		}

		for len(s2) != 0 {
			temp := s2[len(s2)-1] 
			s2 = s2[:len(s2)-1] 
			fmt.Print(temp.data, " ") 

			if temp.left != nil { 
				s1 = append(s1, temp.left) 
			}

			if temp.right != nil {
				s1 = append(s1, temp.right)
			}
		}
	} 
	fmt.Println()
}

func main() {
	root := &BinaryNode{data: 1, left: nil, right: nil}
	root.left = &BinaryNode{data: 2, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 7, left: nil, right: nil}
	root.left.right = &BinaryNode{data: 6, left: nil, right: nil}
	
	root.right = &BinaryNode{data: 3, left: nil, right: nil}
	root.right.left = &BinaryNode{data: 5, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 4, left: nil, right: nil}

	fmt.Println("Spiral Order traversal of binary tree is:")
	printSpiral(root) 
}