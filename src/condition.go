package main

import (
	"fmt"
)

func main() {
	if 1 == 1 {
		fmt.Println("1 is 1")
	} else {
		fmt.Println("1 is not 1")
	}

	n := 3

	switch n {
	case 1:
		fmt.Println("This wont be printed")
	case 3:
		fmt.Println("This will be printed")
		fallthrough
	case 5:
		fmt.Println("This wont be printed without fallthrough keyword 1")
		fallthrough
	case 4:
		fmt.Println("This wont be printed without fallthrough keyword 2")
	default:
		fmt.Println("Default print")
	}
}
