package main

import "fmt"

func main() {
	
	//initializing using type
	var a int = 1   
	var b, c int = 2, 3		
	
	//type can be omitted, if initializer is present
	var d = true   
	var e, f = false, "no!"     
	
	fmt.Println(a, b, c, d, e, f)
	
	// Short variable declarations
	// Outside a function,
	// 		every statement begins with a keyword (var, func, and so on) and
	// 		so the := construct is not available.
	// Inside a function,
	// 		the := short assignment statement can be used in place of a var declaration with implicit type.
	
	// initializer := short assignment => var declaration with implicit type
	g := 1                                 
	h, i, j := true, false, "no!"

	fmt.Println(g, h, i, j)
	
	//var str1 string,  n1 int 				// unexpected comma, expecting semicolon or newline
    //var str2 string = "hello",  n2 int = 1 // unexpected comma, expecting semicolon or newline
	
	// factored var statement
	var (           
		k bool       = false
		l uint64     = 1<<64 - 1
		m string 	 = "hello" 
	)
	
	fmt.Println(k, l, m)

	//Variables declared without an explicit initial value are given their zero value.
	//The zero value is:
	//	0 for numeric types,
	//	false the boolean type, and
	//	"" (the empty string) for strings.
	
	var (
		n bool
		o int
		p string
	)
	
	fmt.Println(n, o, p)

}
