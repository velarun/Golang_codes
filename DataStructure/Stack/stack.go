package main

import "fmt"

type Stack []int

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

func main() {
	s := &Stack{}
	fmt.Println(s.isEmpty())
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(*s)
	fmt.Println(s.pop())
	fmt.Println(*s)
}
