package main

import (
	"fmt"
)

func main() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	var name string
	name = "Shanto"

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(name)

	// multiline varaible declaring
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// multiline variable within a group
	var (
		e int
		f int    = 1
		g string = "hello"
	)

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// cosntant variable
	const (
		A int = 1    // Typed
		B     = 3.14 // Untyped
		C     = "Hi!"
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	// Print Function
	var i, j string = "Hello", "World"

	fmt.Print(i, "\n", j, "\n\n")

	// Println Function
	var s, t string = "Hello", "World"

	fmt.Println(s, t)

	// Printf Function
	var k string = "Hello"
	var l int = 15

	fmt.Printf("i has value: %v and type: %T\n", k, k)
	fmt.Printf("j has value: %v and type: %T", l, l)

	// Formatting Varaibles (Generals)
	var h = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", h)
	fmt.Printf("%#v\n", h)
	fmt.Printf("%v%%\n", h)
	fmt.Printf("%T\n", h)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
}
