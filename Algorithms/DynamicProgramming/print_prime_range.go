//Sieve of Eratosthenes algorithm
package main

import (
	"fmt"
)

const MAX = 1000 

var prefix [MAX+1]int 

func buildPrefix() { 
 
	prime := make([]bool, MAX+1) 
	for i:=0; i<MAX+1; i++ {
		prime[i] = true
	}

	for p := 2; p * p <= MAX; p++ { 
		if prime[p] == true { 
			for i := p * 2; i <= MAX; i += p { 
				prime[i] = false
			}
		} 
	} 

	prefix[0], prefix[1] = 0, 0 
	for p := 2; p <= MAX; p++ { 
		prefix[p] = prefix[p - 1]
		if prime[p] {
			prefix[p]++ 
		} 
	}
} 
 
func query(L, R int) int { 
	return prefix[R] - prefix[L - 1]
} 

func main() { 
	buildPrefix()
	L, R := 5, 10 
	fmt.Println(query(L, R)) 

	L, R = 1, 10 
	fmt.Println(query(L, R)) 
} 