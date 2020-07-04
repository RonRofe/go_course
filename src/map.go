package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Albert":      32,
		"Shoshan Sra": 54,
	}

	fmt.Println(m)
	fmt.Println(m["Albert"])

	v, ok := m["Shoara"]

	fmt.Println(v)
	fmt.Println(ok)
}
