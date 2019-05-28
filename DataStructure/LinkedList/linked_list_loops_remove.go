package main

import "fmt"

type Node struct {
	next *Node
	data int
}

func (list *Node) Display() {
	fmt.Println("\n=====================")
	for list != nil {
		fmt.Print(list.data, "->")
		list = list.next
	}
	fmt.Println("\n=====================")
}

func (l *Node) DetectAndRemoveLoops() {
	if l.next == nil {
		return
	}

	if l.next.next == nil {
		return
	}

	slow := l.next
	fast := l.next

	slow = slow.next
	fast = fast.next.next

	for fast != nil {
		if fast.next == nil {
			break
		}

		if fast == slow {
			break
		}

		slow = slow.next
		fast = fast.next.next
	}

	if slow == fast {
		slow = l.next
		for slow.next != fast.next {
			slow = slow.next
			fast = fast.next
		}
		fast.next = nil
	}

}

func main() {
	list := &Node{data: 6}
	list.next = &Node{data: 5}
	list.next.next = &Node{data: 4}
	list.next.next.next = &Node{data: 3}
	list.next.next.next.next = &Node{data: 2}
	list.next.next.next.next.next = &Node{data: 1}

	list.next.next.next.next.next.next = list.next.next.next

	list.DetectAndRemoveLoops()

	list.Display()
}
