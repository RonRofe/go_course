package main

import (
	"fmt"
)

var y = 42
var z = "test of Z"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}