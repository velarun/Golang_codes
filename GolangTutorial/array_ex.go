// In Go, an _array_ is a numbered sequence of elements of a
// specific length.

package main

import "fmt"

type emp struct {
	code int
	name string
}

func main() {

    // Here we create an array `a` that will hold exactly
    // 5 `int`s. The type of elements and length are both
    // part of the array's type. By default an array is
    // zero-valued, which for `int`s means `0`s.
    var a [5]int
    fmt.Println("Max Array Length without initiate:", a)
	
	// Nil slice length & capacity 0, zero value of a slice is nil
	var z []int                         
    fmt.Println(z, len(z), cap(z))      // [] 0 0
	if z == nil {                       // condition is true
		fmt.Println("nil!")             // nil!
	}

    // We can set a value at an index using the
    // `array[index] = value` syntax, and get a value with
    // `array[index]`.
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
	
	//Setting using append on dynamic array
	z = append(z, 3, 4)
	fmt.Println("set:", z)
	
    // The builtin `len` returns the length of an array.
    fmt.Println("len:", len(a))
	
	// Array Slice
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

    // Use this syntax to declare and initialize an static array(MAX set).
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Static Array:", b)
	
	// Use this syntax to declare and initialize an dynamic array
	c := []string{"g", "h", "i"}
    fmt.Println("Dynamic Array:", c)

    // Array types are one-dimensional, but you can
    // compose types to build multi-dimensional data
    // structures.
	// Static Array
    var twoSD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoSD[i][j] = i + j
        }
    }
    fmt.Println("2d Static: ", twoSD)
	
	// Dynamic Array 2D Array throws panic error 
	//var twoSD [][]int
    //for i := 0; i < 2; i++ {
    //    for j := 0; j < 3; j++ {
    //        twoSD[i][j] = i + j  //--> runtime error: index out of range
    //    }
    //}
	
	// Array initialize using make
	// Dynamically Array can be increased using append
	s1 := make([]int, 2, 4)
	s1[0] = 1
	s1[1] = 2
	fmt.Printf("Initial address and value: %p: %[1]v\n", s1)
	s1 = append(s1, 3, 4)
	fmt.Printf("After first append:        %p: %[1]v\n", s1)
	
	// Static 2D Array using make
	twoD2 := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD2[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD2[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD2)
	
	// append in struct array
    var mySlice []emp
    mySlice = append(mySlice,emp{5,"pizza"})
	fmt.Println("Append in struct array: ", mySlice)
	
	// Without append
	f1 := []emp{emp{1, "burger"}, emp{3, "fanta"}, emp{4, "coke"}}
	fmt.Println("Without append in struct array: ", f1)
	
	// With append
	f2 := append([]emp{}, emp{1, "butter"})
	fmt.Println("With append in struct array: ", f2)
	
	s11 := []int{1, 2, 3, 4}
	fmt.Printf("s11: %p: %[1]v\n", s11)
	
	var s13 []int
	copy(s13, s11)
	fmt.Printf("s13: %p: %[1]v\n", s13)
	
	// Capacity of slice --> cap(s11)
	// way1
	// s12 := make([]int, len(s11), (cap(s11))*2)
	// way2
	s12 := make([]int, 4, 8)
	copy(s12, s11)
	
	fmt.Printf("s12: %p: %[1]v\n", s12)
	s12 = append(s12, 5, 6, 7, 8)
	
	fmt.Printf("s12: %p: %[1]v\n", s12)
	fmt.Printf("s11: %p: %[1]v\n", s11)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
