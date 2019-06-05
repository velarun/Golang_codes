package main

import (
	"fmt"
	"strconv"
)

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

//O(n1 + n2) ->length of strings
func findSum(str1, str2 string) string { 
	
	if len(str1) > len(str2) { 
		temp := str1 
		str1 = str2 
		str2 = temp 
	}

	var str3 string
 
	n1 := len(str1) 
	n2 := len(str2) 
	diff := n2 - n1 

	carry := 0

	for i:=n1-1; i>=0; i-- { 
		sum := int(str1[i]-'0') + int(str2[i+diff]-'0') + carry 
		str3 = fmt.Sprintf("%s%s", str3, strconv.Itoa(sum%10))
		carry = sum/10
	}

	for i:=n2-n1-1; i>=0; i-- { 
		sum := int(str2[i]-'0') + carry
		str3 = fmt.Sprintf("%s%s", str3, strconv.Itoa(sum%10))
		carry = sum/10
	}

	if carry != 0 {
		str3 = fmt.Sprintf("%s%s", str3, string(carry+'0')) 
	}
	
	str3 = reverse(str3) 

	return str3 
}

func isSmaller(str1, str2 string) bool { 

	n1 := len(str1)
	n2 := len(str2) 
  
    if n1 < n2 { 
		return true
	} 

    if n2 > n1 {
		return false
	} 
  
    for i := 0; i < n1; i++ { 
        if str1[i] < str2[i] { 
            return true
		} else if str1[i] > str2[i] {
			return false
		}
	} 
	
    return false
} 

func findDiff(str1, str2 string) string { 
	
	if isSmaller(str1, str2) { 
        t := str1 
        str1 = str2 
        str2 = t 
    } 
  
	var str3 string
  
    n1 := len(str1) 
	n2 := len(str2) 
	diff := n1 - n2 

	carry := 0
  
	for i:=n2-1; i>=0; i-- { 
		sub := int(str1[i + diff] -'0') - int(str2[i]-'0') - carry 
		if sub < 0 { 
            sub = sub + 10 
            carry = 1 
        } else {
            carry = 0 
		}
		str3 = fmt.Sprintf("%s%s", str3, strconv.Itoa(sub))
    } 
		
	for i:=n1-n2-1; i>=0; i-- { 
        if str1[i] == '0' && carry > 0 { 
			str3 = fmt.Sprintf("%s%s", str3, "9") 
            continue
		} 
		
		sub := int(str1[i]-'0') - carry 
        if i > 0 || sub > 0 {
			str3 = fmt.Sprintf("%s%s", str3, strconv.Itoa(sub))
		}

        carry = 0 
    } 
  
    str3 = reverse(str3) 

	return str3 
} 

func findMultiply(str1, str2 string) string { 
	
	n1 := len(str1) 
	n2 := len(str2) 

	if n1 == 0 || n2 == 0 { 
		return "0"
	}

	result := make(map[int]int) 
 
	i_n1 := 0 
	i_n2 := 0 

	for i:=n1-1; i>=0; i-- { 
		carry := 0 
		n1 := str1[i] - '0' 

		i_n2 = 0
			 
		for j:=n2-1; j>=0; j-- { 
			n2 := str2[j] - '0' 
			sum := int(n1) * int(n2) + int(result[i_n1+i_n2]) + int(carry) 
			carry = sum/10 
			result[i_n1 + i_n2] = sum % 10 
			i_n2++ 
		} 

		if carry > 0 {
			result[i_n1 + i_n2] += carry
		}

		i_n1++ 
	} 

	i := len(result) - 1 
	for i>=0 && result[i] == 0 { 
		i-- 
	}

	if i == -1 {
		return "0"
	} 
 
	var s string
	
	for i >= 0 { 
		s = fmt.Sprintf("%s%s", s, strconv.Itoa(result[i])) 
		i--
	}

	return s 
} 

func findDivision(num string, divisor int) string { 

	var quo string
 
	idx := 0 
	temp := int(num[idx] - '0')
	for temp < divisor {
		idx++
		temp = temp * 10 + int(num[idx] - '0')	
	}
	idx++ 
  
	for len(num) > idx {  
		quo = fmt.Sprintf("%s%s", quo, strconv.Itoa(temp/divisor))
		temp = (temp % divisor) * 10 + int(num[idx] - '0') 
		idx++
	} 

	quo = fmt.Sprintf("%s%s", quo, strconv.Itoa(temp/divisor))
 
	if len(quo) == 0 { 
		return "0" 
	}

	return quo 
} 

func sortLargeNumbers(arr []string) { 

    for i:=0; i<len(arr)-1; i++ { 
		left := arr[i] 
		right := arr[i + 1] 

		//If length of left != right, then 
        //return the diff of the length else  
        //use compareTo function to compare values.
        if len(left) > len(right) { 
			arr[i] = right 
			arr[i + 1] = left 
			i -= 2
		//Checking if the length of string are equal means
		//Which number is higher
		} else if len(left) == len(right) {
			for j:=0; j<len(left); j++ {
				if left[j] > right[j] {
					arr[i] = right 
					arr[i + 1] = left
					i-=2
					break
				} else if left[j] == right[j] {
					continue
				} else {
					break
				}
			}
		}
		
		//If array's 0th index has maximum value than 1st means 
		//i value should be -2 now so reducing it to -1
		if i < -1 {
			i = -1
		}
	} 
}  
	
func main() {
	str1 := "919"
	str2 := "198112"
	fmt.Println(findSum(str1, str2)) 
	fmt.Println(findDiff(str1, str2)) 
	fmt.Println(findMultiply(str1, str2)) 
	number := "1248163264128256512" 
	divisor := 125  
	fmt.Println(findDivision(number, divisor))

	arr := []string{"1598078", "1237637463746732323", "97987", "1598078", "1598077"}
	sortLargeNumbers(arr)
	fmt.Println(arr)
}