package main

import "fmt"

type person struct {
    name string
    age  int
}

func main() {

    // This syntax creates a new struct.
    fmt.Println(person{"Apple", 20})

    // You can name the fields when initializing a struct.
    fmt.Println(person{name: "Bat", age: 30})

    // Omitted fields will be zero-valued.
    fmt.Println(person{name: "Cat"})

    // An `&` prefix yields a pointer to the struct.
    fmt.Println(&person{name: "Dog", age: 40})

    // Access struct fields with a dot.
    s := person{name: "Elephant", age: 50}
    fmt.Println(s.name)

    // You can also use dots with struct pointers - the
    // pointers are automatically dereferenced.
    sp := &s
    fmt.Println(sp.age)

    // Structs are mutable.
    sp.age = 51
    fmt.Println(sp.age)
	
	var Person1 person    /* Declare Person1 of type Person */
	var Person2 person    /* Declare Person2 of type Person */

	/* Person 1 specification */
	Person1.name = "Fish"
	Person1.age = 10

	/* Person 2 specification */
	Person2.name = "Giraffe"
	Person2.age = 20
	
	/* Person 2 specification */
	Person3 := person{"Hen", 50}

	/* print Person1 info by use the member access operator (.) */
	fmt.Printf( "Person 1 name : %s\n", Person1.name)
	fmt.Printf( "Person 1 age : %d\n", Person1.age)
	
	//Structures as Function Arguments
	printperson(Person2)
	
	//Pointer Structures as Function Arguments
	printper(&Person3)
}

func printperson( Person person ) {
	fmt.Printf( "Person name : %s\n", Person.name)
	fmt.Printf( "Person age : %d\n", Person.age)
}

func printper( Person *person ) {
	fmt.Printf( "Person name : %s\n", Person.name)
	fmt.Printf( "Person age : %d\n", Person.age)
}