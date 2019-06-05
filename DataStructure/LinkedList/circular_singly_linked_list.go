package main

import "fmt"

type Node struct {
	next *Node
	data int
}

type List struct {
	head *Node
}

func (l *List) Push(node *Node) {

	temp := l.head
	node.next = l.head

	if l.head != nil {
		for temp.next != l.head {
			temp = temp.next
		}
		temp.next = node
	} else {
		node.next = node
	}

	l.head = node
}

func (l *List) Append(node *Node) {

	last := l.head
	node.next = l.head

	if l.head != nil {
		for last.next != l.head {
			last = last.next
		}
		last.next = node 
	} else {
		last = node 
	}
}

func (l *List) Insert(node *Node, prev *Node) { 

	if prev == nil {
		fmt.Println("Prev Node not available")
		return
	}

	node.next = prev.next
	prev.next = node
}

func (l *List) Reverse() { 

  if l.head == nil {  
		return 
	}
	
	current := l.head
	var last *Node
	
	if current != nil {
		next := current.next
		current.next = last
		last = current
		current = next
	}

	for current != l.head {
		next := current.next
		current.next = last
		last = current
		current = next
	}

	current.next = last
	l.head = last
} 

func (l *List) DeleteNode(key int) { 
 
    if l.head == nil {  
		return 
	}

	curr := l.head
	var prev *Node
	var notFound bool

	for curr.data != key { 
		if curr.next == l.head { 
			notFound = true
			break 
		} 

		prev = curr 
		curr = curr.next 
	} 

	if notFound {
		fmt.Println("Given node is not found in the list!!!")
		return
	}

	if curr.next == l.head { 
		l.head = nil 
		return 
	} 

	if curr == l.head { 
		prev = l.head
		for prev.next != l.head {
			prev = prev.next
		}

		l.head = curr.next 
		prev.next = l.head 
	} else if curr.next == l.head { 
		prev.next = l.head 
	} else { 
		prev.next = curr.next
	} 
} 

func (l *List) DeleteNodePos(pos int) {

	temp := l.head

	if temp == nil {
		return
	}

	if pos == 0 {
		prev := l.head
		for prev.next != l.head {
			prev = prev.next
		}

		l.head = temp.next
		prev.next = l.head
		temp = nil
		return
	}

	for i := pos - 1; i > 0; i-- {
		temp = temp.next
		if temp == nil {
			break
		}
	}

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

func (l *List) DeleteList() {
	after := l.head
		
	for after != nil {
		before := after
		after = after.next
		before.next = nil
		before = nil
	}
	
	l.head = nil
	after = nil
}

func (l *List) isCircular() bool {

	if l.head == nil {
		return true
	}

	node := l.head

	for node != nil && node != l.head {
		node = node.next
	}

	return node == l.head
}

func (l *List) printList() {
	temp := l.head
	if l.head != nil {
		for {
			fmt.Printf("%d-->", temp.data)
			temp = temp.next
			if temp == l.head {
				break
			}
		}
	}
	fmt.Println()
}

func main() {
	l := &List{}
	l.Push(&Node{data: 1})
	l.printList()
	l.Push(&Node{data: 2})
	l.printList()
	l.Append(&Node{data: 3})
	l.printList()
	l.Append(&Node{data: 4})
	l.printList()
	l.Insert(&Node{data: 5}, l.head)
	l.printList()
	l.Insert(&Node{data: 6}, l.head.next.next)
	l.printList()
	if l.isCircular() {
		fmt.Println("Circular Linked List")
	} else {
		fmt.Println("Not a Circulr Linked List")
	}
	l.DeleteNode(8)
	l.printList()
	l.DeleteNode(2)
	l.printList()
	l.DeleteNodePos(0)
	l.printList()
	l.DeleteNodePos(1)
	l.printList()
	l.Reverse()
	l.printList()
	l.DeleteList()
	l.printList()
}
