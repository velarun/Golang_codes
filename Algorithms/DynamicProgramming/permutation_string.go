package main

import "fmt"

var res []string

func permute(a []byte, l, r int) {
	if l == r {
		res = append(res, string(a))
	} else {
		for i := l; i < r+1; i++ {
			a[l], a[i] = a[i], a[l]
			permute(a, l+1, r)
			a[l], a[i] = a[i], a[l]
		}
	}
}

func main() {
	str := "ABC"
	n := len(str)
	a := []byte(str)
	permute(a, 0, n-1)
	fmt.Println(res)
}
