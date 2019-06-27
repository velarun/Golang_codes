package main

import "fmt"

type Stack []int

var s1 Stack
var s2 Stack

func (s *Stack) isEmpty() bool {
	return (len(*s) == 0)
}

func (s *Stack) push(item int) {
	*s = append(*s, item)
}

func (s *Stack) pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func enQueue(x int) { 
	for !s1.isEmpty() { 
		s2.push(s1.pop()) 
	} 

	s1.push(x) 

	for !s2.isEmpty() { 
		s1.push(s2.pop()) 
	} 
} 

func deQueue() int { 

	if s1.isEmpty() { 
		fmt.Println("Q is Empty") 
		return 0 
	} 

	x := s1.pop()
	
	return x 
} 

func main() { 
	enQueue(1) 
	enQueue(2) 
	enQueue(3) 

	fmt.Println(deQueue()) 
	fmt.Println(deQueue()) 
	fmt.Println(deQueue()) 
}