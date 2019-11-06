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

func printNGE(arr []int) {
	s := &Stack{} 
	element := 0
	next := 0

	s.push(arr[0])
	
	for i:=1; i<len(arr); i++ { 
		next = arr[i] 
		if s.isEmpty() == false {
			element = s.pop() 
			//If the popped element is smaller than next, then 
			//a) print the pair 
			//b) keep popping while elements are smaller and 
			//stack is not empty
			for element < next { 
				fmt.Println(element, " -- ", next) 
				if s.isEmpty() == true { 
					break
				}
				element = s.pop() 
			}
			//If element is greater than next, then push the element back
			if element > next { 
				s.push(element)
			} 
		}
		//push next to stack so that we can find next greater for it
		s.push(next) 
	}

	//After iterating over the loop, the remaining 
	//elements in stack do not have the next greater 
	//element, so print -1 for them

	for s.isEmpty() == false { 
		element = s.pop() 
		next = -1
		fmt.Println(element, " -- ", next)
	}
}

func main() {
	arr := []int{11, 13, 21, 3} 
	printNGE(arr) 
}