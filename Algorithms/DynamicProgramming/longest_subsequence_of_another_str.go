package main

import (
	"fmt"
	"math"
)

func maxSubsequenceSubstring(x, y string, n, m int) int { 
	
	r := n * m

	dp := make([][]int, r)
	for i:=0; i<r; i++ {
		dp[i] = make([]int, r)
	}

	for i:=1; i<m + 1; i++ { 
		for j:=1; j<n + 1; j++ { 
			fmt.Println(i, j, string(x[j - 1]), string(y[i - 1]), dp[i][j], dp[i - 1][j - 1], dp[i][j - 1])
			if x[j - 1] == y[i - 1] {
				fmt.Println(i-1, j-1) 
				dp[i][j] = 1 + dp[i - 1][j - 1] 
			} else { 
				fmt.Println(i, j-1) 
				dp[i][j] = dp[i][j - 1]
			}
			fmt.Println(dp[i][j])
		}
	} 
				
	ans := 0
	for i:=1; i<m + 1; i++ { 
		ans = int(math.Max(float64(ans), float64(dp[i][n])))
	} 

	return ans 
}

func main() { 
	x := "ABCD"
	y := "BACDBDCD"
	n := len(x) 
	m := len(y) 
	fmt.Println(maxSubsequenceSubstring(x, y, n, m)) 
}