package main

import (
	"fmt"
)

func foo(name string) bool {
	fmt.Println("Hello from foo,", name)
	return true
}

func main() {
	success := foo("Ron")
	fmt.Println(success)
	fmt.Println(multi(1, 1.34))
}

func multi(a int, b float64) (int, float64) {
	return a, b
}
