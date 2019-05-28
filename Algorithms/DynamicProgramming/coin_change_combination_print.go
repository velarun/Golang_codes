package main

import "fmt"

func count(S []int, m int, n int) int { 

 if (n == 0) { 
  return 1
    }
    
 if (n < 0) {
     return 0
 }
  
 if (m <=0 && n >= 1) { 
  return 0
    }
    
 return count( S, m - 1, n ) + count( S, m, n-S[m-1] ) 
}

func main() {
    
    var t int
    var c []int=[]int{1, 3, 5}

    fmt.Scan(&t)
    
    for i:=0; i<t; i++ {
        var n int
        fmt.Scan(&n)

        fmt.Println(count(c, len(c), n))
    }
}
