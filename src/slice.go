package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Martini"}
	fmt.Println(jb)
	mp := []string{"This", "Map", "Three"}
	fmt.Println(jb)

	xp := [][]string{jb, mp}

	fmt.Println(xp)
	// x := []int{1, 3, 4, 5}

	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(x[0])

	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }

	// fmt.Println(x[1:])
	// fmt.Println(x[1:3])

	// x = append(x, 77, 88, 99, 1014)

	// fmt.Println(x)

	// y := []int{0, 0, 0, 0}

	// x = append(x, y...)
	// fmt.Println(x)

	// x = append(x[:2], x[4:]...)
	// fmt.Println(x)

	// x := make([]int, 10, 100)

	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))
}
