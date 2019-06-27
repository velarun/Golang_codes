package main

import "fmt"
import "math"

func stringLen(str string) int { 
	m := make(map[int]int) 
 
	m[0] = -1
	
	count_0 := 0
	count_1 := 0
	res := 0

	for i :=0; i<len(str); i++ { 
		if str[i] == '0' { 
			count_0++
		} else { 
			count_1++
		}

		//If difference between current counts already exists, then 
		//substring between previous and current index has same no. of 
		//0s and 1s. Update result if this substring is more than current result. 
		if _, ok := m[count_1 - count_0]; ok { 
			res = int(math.Max(float64(res), float64(i - m[count_1 - count_0]))) 
		//If current difference is seen first time. 
		} else {
			m[count_1 - count_0] = i
		} 
	}
	return res 
}

func main() {
	str := "101001000"
	fmt.Println("Length of longest balanced sub string = ", stringLen(str)) 
}