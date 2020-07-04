package main

import (
	"fmt"
)

func foo(name string) bool {
	fmt.Println("Hello from foo,", name)
	return true
}

func main() {
	// success := foo("Ron")
	// fmt.Println(success)
	// fmt.Println(multi(1, 1.34))
	unlimited(1, 4, 6, 7, 9, 95)
}

func multi(a int, b float64) (int, float64) {
	return a, b
}

func unlimited(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, val := range x {
		sum += val
	}

	fmt.Println(sum)
}
