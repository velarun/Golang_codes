package main

import "fmt"
import "strings"

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func reverse_words(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func reverse_alphabets(s string) string {
	chars := []rune(s)

    for i, j := 0, len(chars)-1; i<j; {
		if ((chars[i] >= 'a' && chars[i] <= 'z') || 
			(chars[i] >= 'A' && chars[i] <= 'Z')) && 
			((chars[j] >= 'a' && chars[j] <= 'z') || 
			(chars[j] >= 'A' && chars[j] <= 'Z')){
			chars[i], chars[j] = chars[j], chars[i]
			i++
			j--
        } else {
            if !((chars[i] >= 'a' && chars[i] <= 'z') || (chars[i] >= 'A' && chars[i] <= 'Z')) {
                i++
            }
            
            if !((chars[j] >= 'a' && chars[j] <= 'z') || (chars[j] >= 'A' && chars[j] <= 'Z')) {
                j--
            }
		}
	}
    
    return string(chars)
}

func main() {

	a := "hello world!"
	var out string 
	
	for _, v := range a {
		out = string(v) + out
	}

  	fmt.Println("Original String: ", a)
	fmt.Println("Reversed character: ", out)
	
	fmt.Printf("Reverse character using swap: %v\n", reverse("abcdefg"))
	fmt.Println("Reverse words using swap: ", reverse_words("one two three"))

	fmt.Println("Reverse only alphabets using swap: ", reverse_alphabets("@fe&d%cba$"))

}
