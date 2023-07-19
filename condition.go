// package main

// import (
// 	"fmt"
// )

// func main() {
// 	x := 20
// 	y := 18
// 	if x > y {
// 		fmt.Println("x is greater than y")
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	temperature := 14
// 	if temperature > 15 {
// 		fmt.Println("It is warm out there")
// 	} else {
// 		fmt.Println("It is cold out there")
// 	}
// }

package main

import (
	"fmt"
)

func main() {
	time := 22
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
}
