package main

import "fmt"
import s "strings"

func main() {

	a := "hello world!"
	out := 0
	for i, _ := range a {
		out = i
	}

  fmt.Println("Original String: ", a)
	fmt.Println("String Length without len: ", out)
	fmt.Println("String Length with len: ", len(a) - s.Count(a, " "))
	
}
