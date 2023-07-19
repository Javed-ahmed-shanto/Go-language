package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3}   // With var
	arr2 := []int{4, 5, 6, 7, 8} // Using := and without using length

	fmt.Println(arr1)
	fmt.Println(arr2, "\n")

	// String Arrays
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars, "\n")

	// Accessing and changing elements of Arrays
	prices := [3]int{10, 20, 30}
	prices[2] = 50

	fmt.Println(prices[2])

	// Initialize only specific elements
	arr5 := [5]int{1: 10, 2: 40}

	fmt.Println(arr5)

	// Slicing
	arr10 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr10[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	// Appending elements to the end of the slice
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// Creating a copy for memory efficency
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
