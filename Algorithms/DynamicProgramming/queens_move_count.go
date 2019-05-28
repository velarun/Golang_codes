package main

import (
	"fmt"
	"math"
)

func noOfPosition(n int, x int, y int) int {

	// d11, d12, d21, d22 are for diagnoal distances.
	// r1, r2 are for vertical distance.
	// c1, c2 are for horizontal distance.
	var d11, d12, d21, d22, r1, r2, c1, c2 int

	// Initialise the distance to end of the board.
	d11 = int(math.Min(float64(x-1), float64(y-1)))
	d12 = int(math.Min(float64(n-x), float64(n-y)))
	d21 = int(math.Min(float64(n-x), float64(y-1)))
	d22 = int(math.Min(float64(x-1), float64(n-y)))

	r1 = y - 1
	r2 = n - y
	c1 = x - 1
	c2 = n - x

	return d11 + d12 + d21 + d22 + r1 + r2 + c1 + c2
}

func main() {

	// Chessboard size
	var n = 5
	// Queen x position
	var Qposx int = 3
	// Queen y position
	var Qposy int = 4

	fmt.Println(noOfPosition(n, Qposx, Qposy))
}
