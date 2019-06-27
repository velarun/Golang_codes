package main

import "fmt"

type Queue []int

var q1 Queue
var q2 Queue

func (q *Queue) isEmpty() bool {
	return (len(*q) == 0)
}

func (q *Queue) enqueue(item int) {
	*q = append(*q, 0)
	copy((*q)[1:], (*q)[0:])
	(*q)[0] = item
}

func (q *Queue) dequeue() int {
	req := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return req
}

func push(x int) { 
        
    q2.enqueue(x) 
 
    for !q1.isEmpty() { 
        q2.enqueue(q1.dequeue()) 
    } 

    q := q1 
    q1 = q2
    q2 = q 
} 

func pop() int { 
	if q1.isEmpty() { 
		fmt.Println("S is Empty") 
		return 0 
	} 

	x := q1.dequeue() 
	
	return x 
	 
} 

func main() { 
	push(1) 
	push(2) 
	push(3) 

	fmt.Println("POP: ", pop()) 
	fmt.Println("POP: ", pop()) 
	fmt.Println("POP: ", pop()) 
}