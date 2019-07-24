package main

import "fmt"

func Validparentheses(openP, closeP int, String string) {
	if openP == 0 && closeP == 0 {
		fmt.Println(String)
	}

	if openP > closeP {
		return
	}

	if openP > 0 {
		newString := fmt.Sprintf("%s%s", String, "(")
		Validparentheses(openP - 1, closeP, newString)
	}

	if closeP > 0 {
		newString := fmt.Sprintf("%s%s", String, ")")
		Validparentheses(openP, closeP - 1, newString)
	}
}

func printParentheses(n int) {
	Validparentheses(n/2, n/2, "")
}

func main() {
	n := 6
	printParentheses(n)
}