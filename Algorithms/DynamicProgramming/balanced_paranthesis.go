package main

import "fmt"

type Stack []string

func (s *Stack) isEmpty() bool {
	return (len(*s) == 0)
}

func (s *Stack) push(item string) {
	*s = append(*s, item)
}

func (s *Stack) pop() string {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func isMatchingPair(character1, character2 string) bool { 
	if character1 == string("(") && character2 == string(")") { 
		return true
	} else if character1 == string("{") && character2 == string("}") {
		return true
	} else if character1 == string("[") && character2 == string("]") { 
		return true
	} else {
		return false 
	}
} 
	
func areParenthesisBalanced(exp []string) bool { 
	
	st := &Stack{}
	
	/* Traverse the given expression to 
		check matching parenthesis */
	for i:=0; i<len(exp); i++ { 
			
		/*If the exp[i] is a starting 
			parenthesis then push it*/
		if exp[i] == string("{") || exp[i] == string("(") || exp[i] == string("[") {
			st.push(exp[i])
		}
	
		/* If exp[i] is an ending parenthesis 
			then pop from stack and check if the 
			popped parenthesis is a matching pair*/
		if exp[i] == string("}") || exp[i] == string(")") || exp[i] == string("]") { 
			/* If we see an ending parenthesis without 
				a pair then return false*/
			if st.isEmpty() { 
				return false 
			} else if !isMatchingPair(st.pop(), exp[i]) { 
				return false 
			} 
		}	
	} 
		
	/* If there is something left in expression 
		then there is a starting parenthesis without 
		a closing parenthesis */
	if st.isEmpty() {
		return true /*balanced*/
	} else { /*not balanced*/
		return false 
	} 
} 
	
func main() { 
	exp := []string{"{","(","}","[","]"}
	if areParenthesisBalanced(exp) {
		fmt.Println("Balanced ") 
	} else {
		fmt.Println("Not Balanced ")
	} 
} 
