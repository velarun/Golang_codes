package main

import (
	"fmt"
	"strings"
)

func findmatch(mat []string, pat string, x, y, nrow, ncol, level int) bool {

	l := len(pat)

	if level == l {
		return true
	}

	if x < 0 || y < 0 || x >= nrow || y >= ncol {
		return false
	}

	if mat[x][y] == pat[level] {
		temp := string(mat[x][y])
		strings.Replace(mat[x], string(mat[x][y]), "#", 1)

		res := (findmatch(mat, pat, x-1, y, nrow, ncol, level+1) || findmatch(mat, pat, x+1, y, nrow, ncol, level+1) || findmatch(mat, pat, x, y-1, nrow, ncol, level+1) || findmatch(mat, pat, x, y+1, nrow, ncol, level+1))

		strings.Replace(mat[x], string(mat[x][y]), temp, 1)
		return res
	} else {
		return false
	}
}

func checkMatch(mat []string, pat string, nrow, ncol int) bool {

	l := len(pat)

	if l > nrow*ncol {
		return false
	}

	for i := 0; i < nrow; i++ {
		for j := 0; j < ncol; j++ {
			//fmt.Println(string(mat[i][j]))
			if mat[i][j] == pat[0] {
				if findmatch(mat, pat, i, j, nrow, ncol, 0) {
					return true
				}
			}
		}
	}
	return false
}

//4 directions are, Horizontally Left and Right, Vertically Up and Down (no diagonal)
func main() {

	r := 4
	c := 4
	grid := []string{"BACD", "SDES", "DREE", "ADSS"}

	if checkMatch(grid, "ADRESS", r, c) {
		fmt.Println("Word Found")
	} else {
		fmt.Println("Word Not Found")
	}
}
