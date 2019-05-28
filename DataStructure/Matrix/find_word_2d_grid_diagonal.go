package main

import (
	"fmt"
)

var x = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var y = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func search2D(grid [][]byte, row, col int, word string, R, C int) bool {

	if grid[row][col] != word[0] {
		return false
	}

	leng := len(word)

	// Search word in all 8 directions
	// starting from (row,col)
	for dir := 0; dir < 8; dir++ {
		var k int
		rd := row + x[dir]
		cd := col + y[dir]

		// First character is already checked,
		// match remaining characters
		for k = 1; k < leng; k++ {
			// If out of bound break
			if rd >= R || rd < 0 || cd >= C || cd < 0 {
				break
			}

			// If not matched, break
			if grid[rd][cd] != word[k] {
				break
			}

			// Moving in particular direction
			rd += x[dir]
			cd += y[dir]
		}

		// If all character matched, then value of must
		// be equal to length of word
		if k == leng {
			return true
		}
	}
	return false
}

func patternSearch(grid [][]byte, word string, r, c int) {
	for row := 0; row < r; row++ {
		for col := 0; col < c; col++ {
			if search2D(grid, row, col, word, r, c) {
				fmt.Println("pattern found at ", row, ", ", col)
			}
		}
	}
}

//8 directions are, Horizontally Left and Right, Vertically Up, Down and diagonal 4 directions
func main() {
	R := 3
	C := 13
	word := "GEEK"
	grid := [][]byte{[]byte{'G', 'E', 'E', 'K', 'S', 'F', 'O', 'R', 'G', 'E', 'E', 'K', 'S'}, []byte{'G', 'E', 'E', 'K', 'S', 'Q', 'U', 'I', 'Z', 'G', 'E', 'E', 'K'}, []byte{'I', 'D', 'E', 'Q', 'A', 'P', 'R', 'A', 'C', 'T', 'I', 'C', 'E'}}
	patternSearch(grid, word, R, C)
}
