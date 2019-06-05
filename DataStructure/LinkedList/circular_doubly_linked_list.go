package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	data int
}

type List struct {
	head *Node
}

func (l *List) Push(node *Node) {   
   	if l.head == nil {  
      	l.head = node 
      	node.next = l.head   
      	node.prev = l.head  
   	} else {  
       	temp := l.head  
   	 	for temp.next != l.head {  
        	temp = temp.next   
    	}  
		temp.next = node  
		node.prev = temp  
		l.head.prev = node 
		node.next = l.head  
		l.head = node  
   	}   
}  

func (l *List) Append(node *Node) {
 
	if l.head == nil {  
		l.head = node  
		node.next = l.head  
		node.prev = l.head   
	} else {  
		temp := l.head  
		for temp.next != l.head {  
			temp = temp.next  
		}  

		temp.next = node  
		node.prev = temp  
		l.head.prev = node  
		node.next = l.head  
	}  
}  

func (l *List) Insert(node *Node, prev *Node) { 

	if prev == nil {
		fmt.Println("The given previous node cannot be nil")
		return
	}

	next := prev.next
	prev.next = node 
	node.prev = prev 
	node.next = next 
	next.prev = node 
} 

func (l *List) DeleteBeginning() {  

    if l.head == nil {  
        return  
    } else if l.head.next == l.head {  
        l.head = nil      
    } else {  
        temp := l.head   
        for temp.next != l.head {  
            temp = temp.next  
		}  

        temp.next = l.head.next 
		l.head.next.prev = temp 
		l.head = nil   
        l.head = temp.next  
    }  
} 

func (l *List) Deletelast() {  

    if l.head == nil {  
        return  
    } else if l.head.next == l.head {  
        l.head = nil      
    } else {  
        temp := l.head   
        for temp.next != l.head {  
            temp = temp.next  
		}  

        temp.prev.next = l.head  
        l.head.prev = temp.prev    
        temp = nil   
    }  
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
			fmt.Println("Given node is not found in the list!!!")
			notFound = true
			break 
		} 

		prev = curr 
		curr = curr.next 
	}

	if notFound {
		return
	} 

	if curr.next == l.head && prev == nil { 
		l.head = nil 
		return 
	} 

	if curr == l.head { 
		prev = l.head.prev
		l.head = l.head.next 
		prev.next = l.head 
		l.head.prev = prev
	} else if curr.next == l.head { 
		prev.next = l.head 
		l.head.prev = prev
	} else { 
		temp := curr.next
		prev.next = curr.next
		temp.prev = prev
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
		l.head.prev = prev
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
	prev := temp.next.prev
	temp.next = nil
	temp.next = next

	temp.next.prev = prev
}

func (l *List) Reverse() { 

	if l.head == nil {  
		  return 
	  }
	  
	curr := l.head
	var temp *Node
	  
	if curr != nil {
		temp = curr.prev
		curr.prev = curr.next
		curr.next = temp
		curr = curr.prev
	}
  
	for curr != l.head {
		temp = curr.prev
		curr.prev = curr.next
		curr.next = temp
		curr = curr.prev
	}

	if temp != nil {
		l.head = temp.prev
	}
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

func (l *List) printListBackward() {

	last := l.head
	if l.head != nil {
		for last.next != l.head {
			last = last.next
		}

		for last != l.head {
			fmt.Printf("%v <-", last.data)
			last = last.prev
		}
		fmt.Printf("%v <-", last.data)
		fmt.Println()
	}
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
	l.Reverse()
	l.printList()
	l.DeleteBeginning()
	l.printList()
	l.Deletelast()
	l.printList()
	l.DeleteNode(2)
	l.printList()
	l.DeleteNode(5)
	l.printList()
	l.DeleteNodePos(0)
	l.printList()
	if l.isCircular() {
		fmt.Println("Circular Linked List")
	} else {
		fmt.Println("Not a Circulr Linked List")
	}
	l.DeleteNodePos(1)
	l.printList()
	
	l.printListBackward()
	l.DeleteList()
	l.printList()
}
