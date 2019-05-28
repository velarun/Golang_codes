package main

import "fmt"

type Node struct {
	next *Node
	data string
}

type List struct {
	head *Node
}

func (l *List) Push(node *Node) {
	newNode := node
	newNode.next = l.head
	l.head = newNode
}

func (l *List) Append(node *Node) {
	if l.head == nil {
		l.head = node
		return
	}

	last := l.head
	for last.next != nil {
		last = last.next
	}

	last.next = node
}

func (l *List) InsertNodePos(newNode *Node, prevNode *Node) {

	if prevNode == nil {
		fmt.Println("Prev Node not available")
		return
	}

	newNode.next = prevNode.next
	prevNode.next = newNode

}

func (l *List) Reverse() {
	current := l.head
	var last *Node
	for current != nil {
		next := current.next
		current.next = last
		last = current
		current = next
	}

	l.head = last
}

func (l *List) DeleteNodeData(dat string) {

	temp := l.head
	var prev *Node

	if temp != nil {
		if temp.data == dat {
			l.head = temp.next
			temp = nil
			return
		}
	}

	for temp != nil {
		if temp.data == dat {
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
	temp = nil
}

func (l *List) DeleteNodePos(pos int) {

	temp := l.head

	if temp == nil {
		return
	}

	if pos == 0 {
		l.head = temp.next
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
	temp.next = nil
	temp.next = next
}

func (l *List) DeleteNodePointer(node *Node) {

	temp := node.next
	node.data = temp.data
	node.next = temp.next
	temp = nil
}

func (l *List) DeleteList() {
	temp := l.head
	for temp != nil {
		l.head = temp.next
		temp = nil
		temp = l.head
	}
}

func (l *List) Display() {
	fmt.Println("\n=====================")
	list := l.head
	for list != nil {
		fmt.Print(list.data, "->")
		list = list.next
	}
	fmt.Println("\n=====================")
}

func main() {
	l := &List{}
	l.Push(&Node{data: "foo"})
	l.Push(&Node{data: "bar"})
	l.Display()

	l.Append(&Node{data: "car"})
	l.Display()

	l.InsertNodePos(&Node{data: "tar"}, l.head)
	l.InsertNodePos(&Node{data: "par"}, l.head.next.next)
	l.Display()

	l.Reverse()
	l.Display()

	l.DeleteNodeData("foo")
	l.Display()

	l.DeleteNodeData("rar")

	l.DeleteNodePos(0)
	l.Display()

	l.DeleteNodePos(3)
	l.Display()

	l.DeleteNodePointer(l.head)
	l.Display()

	l.DeleteList()
	l.Display()
}
