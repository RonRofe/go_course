package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}

	p2 := person{
		first: "John",
		last:  "Cina",
	}

	fmt.Println(p1.first, p2.last, p2.first, p2.last)
}
