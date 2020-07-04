package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   48,
	}

	p2 := person{
		first: "John",
		last:  "Cina",
		age:   34,
	}

	fmt.Println(p1.first, p2.last, p2.first, p2.last)
}
