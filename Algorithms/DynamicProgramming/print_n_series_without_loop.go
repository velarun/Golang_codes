package main

import "fmt"

func printNumbers(number int){
	
	if number<=0 {
        return
	}

    printNumbers(number-1)

    fmt.Print(number, " ")
}

func main() {
	n := 20
	printNumbers(n)
	fmt.Println()
}