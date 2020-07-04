package main

import (
	"fmt"
)

func main() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Printf("Out: %d\n", i)
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("\t\tIn: %d\n", j)
	// 	}
	// }

	x := 0
	for x < 10 {
		fmt.Printf("%x\t", x)
		x++
	}

	fmt.Println()
}
