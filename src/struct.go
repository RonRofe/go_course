package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   46,
		},
		ltk: false,
	}

	// p2 := person{
	// 	first: "John",
	// 	last:  "Cina",
	// 	age:   34,
	// }

	fmt.Println(sa1.first, sa1.last, sa1.person.age, sa1.ltk)
}
