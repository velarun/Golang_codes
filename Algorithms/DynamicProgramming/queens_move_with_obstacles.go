package main

import (
	"fmt"
	"math"
)

func noOfPosition(n, k, x, y int, obstPosx []int, obstPosy []int) int {

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

	for i := 0; i < k; i++ {
		if x > obstPosx[i] && y > obstPosy[i] && x-obstPosx[i] == y-obstPosy[i] {
			d11 = int(math.Min(float64(d11), float64(x-obstPosx[i]-1)))
		}

		if obstPosx[i] > x && obstPosy[i] > y && obstPosx[i]-x == obstPosy[i]-y {
			d12 = int(math.Min(float64(d12), float64(obstPosx[i]-x-1)))
		}

		if obstPosx[i] > x && y > obstPosy[i] && obstPosx[i]-x == y-obstPosy[i] {
			d21 = int(math.Min(float64(d21), float64(obstPosx[i]-x-1)))
		}

		if x > obstPosx[i] && obstPosy[i] > y && x-obstPosx[i] == obstPosy[i]-y {
			d22 = int(math.Min(float64(d22), float64(x-obstPosx[i]-1)))
		}

		if x == obstPosx[i] && obstPosy[i] < y {
			r1 = int(math.Min(float64(r1), float64(y-obstPosy[i]-1)))
		}

		if x == obstPosx[i] && obstPosy[i] > y {
			r2 = int(math.Min(float64(r2), float64(obstPosy[i]-y-1)))
		}

		if y == obstPosy[i] && obstPosx[i] < x {
			c1 = int(math.Min(float64(c1), float64(x-obstPosx[i]-1)))
		}

		if y == obstPosy[i] && obstPosx[i] > x {
			c2 = int(math.Min(float64(c2), float64(obstPosx[i]-x-1)))
		}
	}

	return d11 + d12 + d21 + d22 + r1 + r2 + c1 + c2
}

func main() {

	// Chessboard size
	var n int = 8
	// No of obstacles
	var k int = 1
	// Queen x position
	var Qposx int = 4
	// Queen y position
	var Qposy int = 4
	// x position of obstacles
	obstPosx := []int{3}
	// y position of obstacles
	obstPosy := []int{5}

	fmt.Println(noOfPosition(n, k, Qposx, Qposy, obstPosx, obstPosy))
}
