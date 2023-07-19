package main

import (
	"fmt"
)

func main() {
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	fruit := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruit {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	for _, val := range fruit {
		fmt.Printf("%v\n", val)
	}
}
