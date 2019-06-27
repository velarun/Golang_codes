package main

import (
	"fmt"
	"math"
)

func longestSubstring(s string) int { 

	cnt := 1 
	maxi := 1
	n := len(s) 

	for i:=1; i<n; i++ { 
		//Check for not consecutive 
		if s[i] != s[i - 1] { 
			cnt++ 
		} else { 
			maxi = int(math.Max(float64(cnt), float64(maxi))) 
			cnt = 1 
		}
	}

	maxi = int(math.Max(float64(cnt), float64(maxi))) 

	return maxi 
}

func main() { 
	s := "ccccdeededff" 
	fmt.Println(longestSubstring(s))
}