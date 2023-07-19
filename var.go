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
}
