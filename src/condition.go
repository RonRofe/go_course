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

	switch {
	case false:
		fmt.Println("This wont be printed")
	case 1 == 1:
		fmt.Println("This will be printed")
		fallthrough
	case 1 == 5:
		fmt.Println("This wont be printed without fallthrough keyword 1")
		fallthrough
	case 3 == 1:
		fmt.Println("This wont be printed without fallthrough keyword 2")
	default:
		fmt.Println("Default print")
	}
}
