package main

import (
	"fmt"
)

var y = 42
var z = "test of Z"
var a = `James said

"Hello world"`

var defStr string
var defInt int

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println("Initial value of string: ", defStr, "!")
	fmt.Println("Initial value of integer: ", defInt, "!")
}
