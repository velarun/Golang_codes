package main

import (
	"fmt"
)

func WordBreak(set map[string]bool, word string) bool {

	size := len(word)

	if size == 0 {
		return true
	}

	for i := 1; i <= size; i++ {
		_, found := set[word[0:i]]
		if found && WordBreak(set, word[i:size]) {
			return true
		}
	}
	return false
}

func main() {
	temp_dictionary := []string{"mobile", "samsung", "sam", "sung", "man", "mango", "icecream", "and", "go", "i", "like", "ice", "cream"}
	set := make(map[string]bool)
	for _, v := range temp_dictionary {
		set[v] = true
	}
	fmt.Println(set)
	delete(set, "mobile")
	fmt.Println(set)
	fmt.Println(WordBreak(set, "ilikesamsung"))
	fmt.Println(WordBreak(set, "nope"))
}
