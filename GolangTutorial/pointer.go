package main

import "fmt" 

func main() {

	var ptr *int //zero value means nil
	
	fmt.Printf("The value of ptr is : %x\n", ptr)
	//fmt.Println("The value of *p before pointing to variable: ", *p) 
	//panic: runtime error: invalid memory address or nil pointer dereference
	//[signal 0xc0000005 code=0x0 addr=0x0 pc=0x48df72]
	
	if (ptr == nil) {
		fmt.Println("Success")
	}
	
	var a int = 20   /* actual variable declaration */
	var p *int      /* pointer variable declaration */ 

	p = &a  /* store address of a in pointer variable*/ 
	//p is pointer to its operand (&) i.e a

	fmt.Printf("Address of a variable: %x\n", &a  )

	/* address stored in pointer variable */
	fmt.Printf("Address stored in p variable: %x\n", p )

	/* read a through the pointer p, pointer's underlying value (*) "dereferencing" or "indirecting */
	fmt.Printf("Value of *p variable: %d\n", *p )
		
	*p = 21             // set i through the pointer p

	// p = p + 1           
	// invalid operation: p + 1 (mismatched types *int and int)     
	// Unlike C, Go has no pointer arithmetic.
	
	fmt.Printf("Value of *p variable after setting through pointer: %d\n", *p )
	
	a=22
	
	fmt.Printf("Value of *p variable after setting through variable: %d\n", *p )
	
	const MAX int = 3
	
	b := []int{10,100,200}
	var i int
	var ptr [MAX]*int;

	for  i = 0; i < MAX; i++ {
	  ptr[i] = &b[i] /* assign the address of integer. */
	}
	
	for  i = 0; i < MAX; i++ {
	  fmt.Printf("Value of b[%d] = %d\n", i,*ptr[i] )
	}
	
	var c int
	var ptr1 *int
	var pptr **int

	a = 3000

	/* take the address of var */
	ptr1 = &a

	/* take the address of ptr using address of operator & */
	pptr = &ptr1

	/* take the value using pptr */
	fmt.Printf("Value of a = %d\n", a )
	fmt.Printf("Value available at *ptr = %d\n", *ptr1 )
	fmt.Printf("Value available at **pptr = %d\n", **pptr)
	
	
}