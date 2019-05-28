package main

import (
	"fmt"
)

func CheckPalindromicSubString(str string) {

	maxlength := 1
	low := 0
	high := 0
	start := 0

	length := len(str)

	for i := 1; i < length; i++ {
		//Even possibilites of palindrome
		low = i - 1
		high = i
		for low >= 0 && high < length && str[low] == str[high] {
			if high-low+1 > maxlength {
				start = low
				maxlength = high - low + 1
			}
			low--
			high++
		}

		//Odd Possibilities of palindrome
		low = i - 1
		high = i + 1
		for low >= 0 && high < length && str[low] == str[high] {
			if high-low+1 > maxlength {
				start = low
				maxlength = high - low + 1
			}
			low--
			high++
		}
	}

	fmt.Println("Longest palindromic length from string:", maxlength)
	fmt.Println("Palindromic substring from string:", str[start:start+maxlength])
}

func main() {
	str := "tooharddrahootfreakout"
	CheckPalindromicSubString(str)
}
