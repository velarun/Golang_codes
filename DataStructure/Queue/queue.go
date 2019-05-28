package main

import "fmt"

type Queue []int

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

func main() {
	q := &Queue{}
	fmt.Println(q.isEmpty())
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	fmt.Println(*q)
	fmt.Println(q.dequeue())
	fmt.Println(*q)
}
