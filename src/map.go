package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Albert":      32,
		"Shoshan Sra": 54,
	}

	// fmt.Println(m)
	// fmt.Println(m["Albert"])

	// v, ok := m["Shoara"]

	// fmt.Println(v)
	// fmt.Println(ok)

	// if v, ok := m["Albert"]; ok {
	// 	fmt.Println("THIS IS THE IF PRINT", v)
	// }

	m["ron"] = 23

	for k, v := range m {
		fmt.Println(k, v)
	}
}
