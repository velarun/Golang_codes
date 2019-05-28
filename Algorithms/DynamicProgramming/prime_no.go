package main

import (
	"fmt"
	"math"
)

func IsPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}

func main() {
    
    fmt.Println("Printing sequence of Prime number between 1 to 100:")
    for i := 1; i <= 100; i++ {
        if IsPrime(i) {
            fmt.Printf("%v ", i)
	}
    }
    
    fmt.Println("\nChecking if number is Prime or Not")
    primeNo := 45
    if IsPrime(primeNo) {
	fmt.Printf("%d is a Prime Number\n", primeNo)
    } else {
	fmt.Printf("%d is not a Prime Number\n", primeNo)
    }
}
