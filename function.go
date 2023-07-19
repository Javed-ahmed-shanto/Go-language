package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("I just got executed!\n")
}

func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	myMessage()
	myMessage()
	myMessage()

	// Function with Parameter
	familyName("Liam", 3)
	familyName("Jenny", 14)
	familyName("Anja", 30)

	// Return
	fmt.Println(myFunction(5, "Hello"))

	//Recursion
	fmt.Println(factorial_recursion(4))
}
