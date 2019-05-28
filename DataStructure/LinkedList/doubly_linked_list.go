package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type list struct {
	head *Node
}

func (l *list) Push(key interface{}) {
	new_node := &Node{key: key}

	new_node.next = l.head

	if l.head != nil {
		l.head.prev = new_node
	}

	l.head = new_node
}

func (l *list) Append(key interface{}) {
	new_node := &Node{key: key}
	new_node.next = nil

	if l.head == nil {
		new_node.prev = nil
		l.head = new_node
		return
	}

	last := l.head
	for last.next != nil {
		last = last.next
	}

	last.next = new_node

	new_node.prev = last
}

func (l *list) Insert(key interface{}, prev_node *Node) {

	if prev_node == nil {
		fmt.Println("The given previous node cannot be nil")
		return
	}

	new_node := &Node{key: key}

	new_node.next = prev_node.next

	prev_node.next = new_node

	new_node.prev = prev_node

	if new_node.next != nil {
		new_node.next.prev = new_node
	}
}

func (l *list) DeleteNodeData(dat interface{}) {

	temp := l.head
	var prev *Node

	if temp != nil {
		if temp.key == dat {
			l.head = temp.next
			l.head.prev = nil
			temp = nil
			return
		}
	}

	for temp != nil {
		if temp.key == dat {
			break
		}
		prev = temp
		temp = temp.next
	}

	//if key not found
	if temp == nil {
		fmt.Println("Key not found")
		return
	}

	prev.next = temp.next

	if temp.next != nil {
		temp.next.prev = prev
	}

	temp = nil
}

func (l *list) DeleteNodePos(pos int) {

	temp := l.head

	if temp == nil {
		return
	}

	if pos == 0 {
		l.head = temp.next
		l.head.prev = nil
		temp = nil
		return
	}

	for i := pos - 1; i > 0; i-- {
		temp = temp.next
		if temp == nil {
			break
		}
	}

	//if pos not found
	if temp == nil {
		fmt.Println("Position not found")
		return
	}

	if temp.next == nil {
		fmt.Println("Position not found")
		return
	}

	next := temp.next.next
	prev := temp.next.prev
	temp.next = nil
	temp.next = next

	temp.next.prev = prev
}

func (l *list) DeleteList() {
	temp := l.head
	for temp != nil {
		l.head = temp.next
		temp = nil
		temp = l.head
	}
}

func (l *list) Display() {
	temp := l.head
	for temp != nil {
		fmt.Printf("%+v ->", temp.key)
		temp = temp.next
	}
	fmt.Println()
}

func (l *list) BackwardDisplay() {

	last := l.head
	for last.next != nil {
		last = last.next
	}

	for last != nil {
		fmt.Printf("%v <-", last.key)
		last = last.prev
	}
	fmt.Println()
}

func (l *list) Reverse() {
	curr := l.head
	var temp *Node

	for curr != nil {
		temp = curr.prev
		curr.prev = curr.next
		curr.next = temp
		curr = curr.prev
	}

	if temp != nil {
		l.head = temp.prev
	}
}

func main() {
	link := &list{}
	link.Push(5)
	link.Push(9)

	link.Append(22)
	link.Append(27)

	link.Insert(8, link.head.next)
	link.Insert(56, link.head.next.next)

	fmt.Println("\n==============================\n")
	fmt.Println("Doubly linked list display using next pointer:\n")
	link.Display()
	fmt.Println("\n==============================\n")
	fmt.Println("Doubly linked list display using prev pointer:\n")
	link.BackwardDisplay()
	fmt.Println("\n==============================\n")
	fmt.Println("Doubly linked list reverse display:\n")
	link.Reverse()
	link.Display()
	fmt.Println("\n==============================\n")
	fmt.Println("Doubly linked list Delete data 5:\n")
	link.DeleteNodeData(5)
	link.Display()
	link.BackwardDisplay()
	fmt.Println("\n==============================\n")
	fmt.Println("Doubly linked list Delete data 27 head position:\n")
	link.DeleteNodeData(27)
	link.Display()
	link.BackwardDisplay()
	fmt.Println("\n==============================\n")
	fmt.Println("Doubly linked list Delete data at position 1:\n")
	link.DeleteNodePos(1)
	link.Display()
	link.BackwardDisplay()
	fmt.Println("\n==============================\n")
	fmt.Println("Delete Doubly linked list:\n")
	link.DeleteList()
	link.Display()
	fmt.Println("\n==============================\n")
}
