package main

import (
	"fmt"
	"math"
)

func lenoflongestnonpalindrome(s string) int {

	max1, length := 1, 0

	for i:=0; i<len(s)-1; i++ { 
		//checking palindrome of size 2 example: aa 
		if s[i] == s[i + 1] { 
			length = 0
		//checking palindrome of size 3 example: aba 
		} else if i > 0 && s[i + 1] == s[i - 1] { 
			length = 1
		} else { 
			length++
		}

		max1 = int(math.Max(float64(max1), float64(length + 1))) 
	}

	if max1 == 1 { 
		return 0
	} else { 
		return max1 
	}
}

func main() {
	s := "synapse"
	fmt.Println(lenoflongestnonpalindrome(s)) 
}