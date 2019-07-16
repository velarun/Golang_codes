package main

import "fmt"

func main() {
	char_map := []string{" ", ".,'?", "ABC", "DEF", "GHI", "JKL", "MNO", "PQRS", "TUV", "WXYZ"}
	n := 234
	print_digits(char_map, n, "")
}

func print_digits(char_map []string, n int, suffix string) { 
    if n <= 0 { 
        fmt.Println(suffix) 
	} else {
        for _, c := range char_map[n%10] {
			suffix := fmt.Sprintf("%s%s", string(c), string(suffix)) 
			print_digits(char_map, n/10, suffix) 
		}
	}
}