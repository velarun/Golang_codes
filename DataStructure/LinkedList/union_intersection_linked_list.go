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

func getIntersection(head1, head2 *Node) *List { 
	hset := make(map[int]int)
	n1 := head1 
	n2 := head2
	result := &List{}
	
	// loop stores all the elements of list1 in hset 
	for n1 != nil { 
		if _, ok := hset[n1.data]; !ok {  
			hset[n1.data] = 1
		} 
		n1 = n1.next 
	} 
	
	//For every element of list2 present in hset 
	//loop inserts the element into the result 
	for n2 != nil { 
		if _, ok := hset[n2.data]; ok { 
			result.Push(n2.data)
		} 
		n2 = n2.next 
	} 

	return result
} 

func getUnion(head1, head2 *Node) *List { 
	hset := make(map[int]int)
	n1 := head1 
	n2 := head2
	result := &List{}
	
	// loop stores all the elements of list1 in hset 
	for n1 != nil { 
		if _, ok := hset[n1.data]; !ok {  
			hset[n1.data] = 1
		} else {
			hset[n1.data] += 1
		}
		n1 = n1.next 
	} 

	// loop stores all the elements of list2 in hset 
	for n2 != nil { 
		if _, ok := hset[n2.data]; !ok {  
			hset[n2.data] = 1
		} else {
			hset[n2.data] += 1
		}
		n2 = n2.next 
	}

	for val := range hset {
		result.Push(val)
	} 

	return result
} 

func main() {
	llist1 := &List{}
	llist1.Push(20) 
	llist1.Push(4)
	llist1.Push(15)
	llist1.Push(10)
	llist1.Display()

	llist2 := &List{}
	llist2.Push(10) 
	llist2.Push(2) 
	llist2.Push(4) 
	llist2.Push(8) 
	llist2.Display()

	intersection := getIntersection(llist1.head, llist2.head)
	
	fmt.Println("Intersection List is") 
	intersection.Display()

	union := getUnion(llist1.head, llist2.head)

	fmt.Println("Union List is")
	union.Display() 
} 