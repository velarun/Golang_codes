package main

import "fmt"

type Node struct {
	next *Node
	data int
}

func Traverse(list1 *Node, list2 *Node, dist int) int {
    curr1 := list1
    curr2 := list2
    
    for i:=0; i<dist; i++ {
        if curr1 == nil {
            return -1
        }
        curr1 = curr1.next
    }
    
    for curr1 != nil && curr2 != nil {
        if curr1.data == curr2.data {
            return curr1.data
        }
        curr1 = curr1.next
        curr2 = curr2.next
    }
    
    return -1
}
    
func IntersectionPoint(list1 *Node, list2 *Node) {
	
	curr := list1
    count1 := 0
    
    for curr != nil {
        count1++
        curr = curr.next
    }
    
    curr = list2
    count2 := 0
    
    for curr != nil {
        count2++
        curr = curr.next
    }
  
    if count1 > count2 {
        val := count1 - count2 
        fmt.Println(Traverse(list1, list2, val))
    } else {
        val := count2 - count1 
        fmt.Println(Traverse(list2, list1, val))
    }
}

func (list *Node) Display() {
	fmt.Println("\n=====================")
	for list != nil {
		fmt.Print(list.data, "->")
		list = list.next
	}
	fmt.Println("\n=====================")
}

func main() {
	list1 := &Node{data: 3}
	list1.next = &Node{data: 6}
	list1.next.next = &Node{data: 15}
	list1.next.next.next = &Node{data: 15}
	list1.next.next.next.next = &Node{data: 30}

	list2 := &Node{data: 10}
	list2.next = &Node{data: 15}
	list2.next.next = &Node{data: 30}

	list1.Display()
	list2.Display()

	IntersectionPoint(list1, list2)
}