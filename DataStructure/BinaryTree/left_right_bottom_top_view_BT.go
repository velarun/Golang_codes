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

func leftViewUtil(root *BinaryNode, level int, max_level []int) {

	if root == nil {
		return
	}

	if max_level[0] < level {
		fmt.Printf("%d ", root.data)
		max_level[0] = level
	}

	leftViewUtil(root.left, level+1, max_level)
	leftViewUtil(root.right, level+1, max_level)
}

func leftView(root *BinaryNode) {
	max_level := []int{0}
	leftViewUtil(root, 1, max_level)
	fmt.Println()
}

func rightViewUtil(root *BinaryNode, level int, max_level []int) {

	if root == nil {
		return
	}

	if max_level[0] < level {
		fmt.Printf("%d ", root.data)
		max_level[0] = level
	}

	rightViewUtil(root.right, level+1, max_level)
	rightViewUtil(root.left, level+1, max_level)
}

func rightView(root *BinaryNode) {
	max_level := []int{0}
	rightViewUtil(root, 1, max_level)
	fmt.Println()
}

// structure of pair
type Queue struct {
	first  *BinaryNode
	second int
}

func topView(root *BinaryNode) {

	if root == nil {
		return
	}

	// Take a temporary node
	temp := &BinaryNode{}

	// Queue to do BFS
	q := []Queue{}

	// map to store node at each vertical distance
	mp := make(map[int]*BinaryNode)

	q = append(q, Queue{first: root, second: 0})

	// BFS
	for len(q) > 0 {

		temp = q[0].first
		d := q[0].second
		q = q[1:]

		// If any node is not at that vertical distance
		// just insert that node in map and print it
		if mp[d] == nil {
			fmt.Print(temp.data, " ")
			mp[d] = temp
		}

		// Continue for left node
		if temp.left != nil {
			q = append(q, Queue{first: temp.left, second: d - 1})
		}

		// Continue for right node
		if temp.right != nil {
			q = append(q, Queue{first: temp.right, second: d + 1})
		}
	}
	fmt.Println()
}

func bottomView(root *BinaryNode) {

	if root == nil {
		return
	}

	// Take a temporary node
	temp := &BinaryNode{}

	// Queue to do BFS
	q := []Queue{}

	// map to store node at each vertical distance
	mp := make(map[int]*BinaryNode)

	q = append(q, Queue{first: root, second: 0})

	for len(q) > 0 {

		temp = q[0].first
		d := q[0].second
		q = q[1:]

		mp[d] = temp

		// Continue for left node
		if temp.left != nil {
			q = append(q, Queue{first: temp.left, second: d - 1})
		}

		// Continue for right node
		if temp.right != nil {
			q = append(q, Queue{first: temp.right, second: d + 1})
		}
	}

	for i := range mp {
		fmt.Print(mp[i].data, " ")
	}
	fmt.Println()

}

func main() {
	root := &BinaryNode{data: 20, left: nil, right: nil}
	root.left = &BinaryNode{data: 8, left: nil, right: nil}
	root.left.left = &BinaryNode{data: 5, left: nil, right: nil}
	root.left.right = &BinaryNode{data: 3, left: nil, right: nil}
	root.left.right.left = &BinaryNode{data: 10, left: nil, right: nil}
	root.left.right.right = &BinaryNode{data: 14, left: nil, right: nil}

	root.right = &BinaryNode{data: 22, left: nil, right: nil}
	root.right.left = &BinaryNode{data: 4, left: nil, right: nil}
	root.right.right = &BinaryNode{data: 25, left: nil, right: nil}

	fmt.Println("Left View of Binary Tree:")
	leftView(root)

	fmt.Println("Right View of Binary Tree:")
	rightView(root)

	fmt.Println("Top View of Binary Tree:")
	topView(root)

	fmt.Println("Bottom View of Binary Tree:")
	bottomView(root)
}
