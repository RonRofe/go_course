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
}
