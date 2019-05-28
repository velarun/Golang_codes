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

func printAncestors(root *BinaryNode, key int) {

	if root == nil {
		return
	}

	//Create a stack to hold ancestors
	st := []*BinaryNode{}

	//Traverse the complete tree in postorder way till we find the key
	for {

		//Traverse the left side. While traversing, push
		//the nodes into the stack so that their right
		//subtrees can be traversed later
		for root != nil && root.data != key {
			st = append(st, root)
			root = root.left
		}

		//If the node whose ancestors are to be printed
		//is found, then break the while loop.
		if root != nil && root.data == key {
			break
		}

		//Check if right sub-tree exists for the node at top
		//If not then pop that node because we don't need
		//this node any more.
		if st[len(st)-1].right == nil {
			root = st[len(st)-1]
			st = st[:len(st)-1]

			//If the popped node is right child of top,
			//then remove the top as well. Left child of
			//the top must have processed before.
			for len(st) != 0 && st[len(st)-1].right == root {
				root = st[len(st)-1]
				st = st[:len(st)-1]
			}
		}

		//if stack is not empty then simply set the root
		//as right child of top and start traversing right sub-tree.
		if len(st) == 0 {
			root = nil
		} else {
			root = st[len(st)-1].right
		}
	}

	//If stack is not empty, print contents of stack
	//Here assumption is that the key is there in tree
	for len(st) != 0 {
		fmt.Print(st[len(st)-1].data, " ")
		st = st[:len(st)-1]
	}

	fmt.Println()
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

	key := 7
	printAncestors(root, key)
}
