package main

import (
	"fmt"
)

var y = 42

type hotdog int

var b hotdog

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	y = 911
	fmt.Printf("%#x\t%b\n", y, y)

	s := fmt.Sprintf("This is y binary decoded: %b", y)
	fmt.Println(s)

	fmt.Println("------------- Custom Types -------------")

	b = 43

	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
