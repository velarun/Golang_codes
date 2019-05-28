package main

import (
	"fmt"
	"math"
)

type Empty interface {
	A() float64
}

type Mystring string

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Object interface {
	Volume() float64
}

type Material interface {
	Shape
	Object
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Cube struct {
	side float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return 2 * math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Perimeter() float64 {
	return 12 * c.side
}

func explain(i interface{}) {
	fmt.Printf("Type is %T, Value is %v\n", i, i)
	switch i.(type) {
	case string:
		fmt.Printf("String content %s\n", i)
	default:
		fmt.Println("Type is not in the List")
	}
}

func main() {
	var e Empty

	//fmt.Println(s.Area())
	//panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println("Value of e is", e)
	fmt.Printf("Type of e is %T\n", e)

	//if perimeter func is not defined means it will throw error as
	//cannot use Rect literal (type Rect) as type Shape in assignment:
	//Rect does not implement Shape (missing Perimeter method)
	//In order to successfully implement an interface,
	//you need to implement all methods declared by the interface.

	var s Shape = Rect{5, 6}
	fmt.Println("Area of Rect is", s.Area())
	fmt.Println("Perimeter of Rect is", s.Perimeter())
	fmt.Printf("Type of s is %T\n", s)
	fmt.Printf("Value of s is %v\n", s)

	//var s Shape = Circle{5}
	//throws error as s redeclared in this block

	s = Circle{5}
	fmt.Println("Area of Rect is", s.Area())
	fmt.Println("Perimeter of Rect is", s.Perimeter())
	fmt.Printf("Type of s is %T\n", s)
	fmt.Printf("Value of s is %v\n", s)

	//When an interface has zero methods, it called as empty interface.
	//This is represented by interface{}.
	//Since empty interface has zero methods, all types implement this interface.
	ms := Mystring("Hello World")
	explain(s)
	explain(ms)
	explain("Hello Golang Interface")

	cu := Cube{5}
	var multi1 Shape = cu
	var multi2 Object = cu
	//an interface can not implement other interfaces or extend them,
	//but we can create new interface by merging two or more interfaces
	var m Material = cu

	fmt.Println("Area of Cube is", multi1.Area())
	fmt.Println("Volume of Cube is", multi2.Volume())
	//multi1.Volume undefined (type Shape has no field or method Volume)
	//fmt.Println("Area of Rect is", multi1.Volume())
	//To resolve this we can use embedded interface by merging two interface
	//on another interface

	fmt.Println("Volume of Cube using embedded interface is", m.Volume())
	fmt.Println("Area of Cube using embedded interface is", m.Area())

	//Type Assertion
	//i.(Type) where i is an interface and Type is a type that implements the interface i.
	//Go will check if dynamic type of i is identical to Type.
	var ss Shape = Cube{5}
	c := ss.(Cube)

	fmt.Println("Area of Cube using Assertion is", c.Area())
	fmt.Println("Volume of Cube using Assertion is", c.Volume())

	d := ss.(Object)
	fmt.Println("Volume of Cube using Assertion is", d.Volume())

}
