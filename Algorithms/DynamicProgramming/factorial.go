package main

import "fmt"

func fact(n int) int {	
	f := 1
	for n>=2 {
		f = f * n
		n = n - 1
	}
	return f
}

func recfact(x int) int {
    if x == 0 {
        return 1
    }
    return x * recfact(x - 1)
}

func main() {
	x := 5
	fmt.Println("Factorial: ", fact(x))
	fmt.Println("Recursive Factorial: ", recfact(x))
}

