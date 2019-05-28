package main

import "fmt"

func stringInSlice(a string, list map[string]int) bool {
	for _, b := range list {
		if string(b) == a {
			return true
		}
	}
	return false
}

func findLongestSubstring(str string) string {

	n := len(str)
	st := 0
	maxlen := 0
	start := 0
	pos := make(map[string]int)

	pos[string(str[0])] = 0

	for i := 1; i < n; i++ {
		if stringInSlice(string(str[i]), pos) {
			pos[string(str[i])] = i
		} else {
			if pos[string(str[i])] >= st {
				currlen := i - st
				if maxlen < currlen {
					maxlen = currlen
					start = st
				}
				st = pos[string(str[i])] + 1
			}
			pos[string(str[i])] = i
		}

		if maxlen < i-st {
			maxlen = i - st
			start = st
		}
	}

	return str[start : start+maxlen]
}

func main() {
	str := "GEEKSFORGEEKS"
	fmt.Println(findLongestSubstring(str))
}
