package main

import "fmt"

type Node struct {
	next *Node
	data int
}

type List struct {
	head *Node
}

func (l *List) Push(data int) {
	newNode := &Node{data: data}
	newNode.next = l.head
	l.head = newNode
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

func (l *List) addTwoLists(first, second *Node) { 

	var prev, temp *Node
	carry := 0
	var sdata, fdata int

	for first != nil || second != nil { 
		if first == nil {
			fdata = 0 
		} else {
			fdata = first.data 
		}

		 
		if second == nil {
			sdata = 0
		} else {
			sdata = second.data 
		}

		Sum := carry + fdata + sdata 

		//update carry for next calculation 
		 
		if Sum >= 10 {
			carry = 1
		} else {
			carry = 0
		} 

		//update sum if it is greater than 10 
		
		if Sum < 10 {
			Sum = Sum 
		} else {
			Sum = Sum % 10
		}

		//Create a new node with sum as data 
		temp := &Node{data: Sum}

		if l.head == nil { 
			l.head = temp 
		} else {
			prev.next = temp 
		}

		prev = temp 

		if first != nil { 
			first = first.next
		}

		if second != nil { 
			second = second.next
		}
	}

	if carry > 0 { 
		temp.next = &Node{data: carry}
	}
}

func main() {
	llist1 := &List{}
	llist1.Push(6) 
	llist1.Push(4)
	llist1.Push(9)
	llist1.Push(5)
	llist1.Push(7)
	llist1.Display()

	llist2 := &List{}
	llist2.Push(4)
	llist2.Push(8) 
	llist2.Display()

	res := &List{}
	res.addTwoLists(llist1.head, llist2.head)

	fmt.Println("Resultant list is") 
	res.Display() 
}