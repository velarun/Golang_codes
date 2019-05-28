package main

import "fmt"

type Node struct {
	next   *Node
	random *Node
	data   int
}

func (list *Node) Display() {

	fmt.Println("\n=====================")
	for list != nil {
		fmt.Println("Data:", list.data, ", Random Data:", list.random.data)
		list = list.next
	}
	fmt.Println("\n=====================")
}

func (list *Node) Clone() *Node {
	//Clone a doubly linked list with random pointer
	//with O(1) extra space

	//Insert additional node after every node of original list
	curr := list
	for curr != nil {
		new := &Node{data: curr.data}
		new.next = curr.next
		curr.next = new
		curr = curr.next.next
	}

	//Adjust the random pointers of the newly added nodes
	curr = list
	for curr != nil {
		curr.next.random = curr.random.next
		curr = curr.next.next
	}

	//Detach original and duplicate list from each other
	curr = list
	dup_root := list.next
	for curr.next != nil {
		tmp := curr.next
		curr.next = curr.next.next
		curr = tmp
	}

	return dup_root
}

func main() {
	originalList := &Node{data: 1}
	originalList.next = &Node{data: 2}
	originalList.next.next = &Node{data: 3}
	originalList.next.next.next = &Node{data: 4}
	originalList.next.next.next.next = &Node{data: 5}

	// 1's random points to 3
	originalList.random = originalList.next.next
	// 2's random points to 1
	originalList.next.random = originalList
	// 3's random points to 5
	originalList.next.next.random = originalList.next.next.next.next
	// 4's random points to 5
	originalList.next.next.next.random = originalList.next.next.next.next
	// 5's random points to 2
	originalList.next.next.next.next.random = originalList.next

	fmt.Println("Original list:")
	originalList.Display()

	clonedList := originalList.Clone()

	fmt.Println("Cloned list:")
	clonedList.Display()
}
