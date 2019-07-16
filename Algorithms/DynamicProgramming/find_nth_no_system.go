package main

import "fmt"

func find(n int) { 
	arr := make([]string, n+1) 
	size := 1
	m := 1 

	for size <= n {
		for i:=0; i < m && (size + i) <= n; i++ { 
			arr[size + i] = fmt.Sprintf("%s%s", "3", arr[size - m + i]) 
		}
 
		for i:=0; i < m && (size + m + i) <= n; i++ { 
			arr[size + m + i] = fmt.Sprintf("%s%s", "4", arr[size - m + i])
		}

		m = m*2  
		size = size + m
	}

	fmt.Println(arr[n]) 
}

func main() { 
	l := 16
	for i:=1; i<=l; i++ { 
		find(i)
	}
}
