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

	// x := 0
	// for x < 10 {
	// 	fmt.Printf("%x\t", x)
	// 	x++
	// }

	// fmt.Println()

	// x := 0
	// for {
	// 	if x >= 10 {
	// 		break
	// 	}
	// 	fmt.Printf("%x\t", x)
	// 	x++
	// }

	// fmt.Println()

	x := 0
	for {
		x++
		if x > 100 {
			break;
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)
	}
}
