package main

import (
	"fmt"
)

func main() {
	x := []int{1, 3, 4, 5}

	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(x[0])

	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }

	fmt.Println(x[1:])
	fmt.Println(x[1:3])
}
