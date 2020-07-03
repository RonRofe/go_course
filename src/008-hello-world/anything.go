package main

import "fmt"

func main() {
	fmt.Println("Hello class! test string")
	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("Second string")

	bar()
}

func bar() {
	fmt.Println("Third string")
}
