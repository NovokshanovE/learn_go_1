package main

import "fmt"

func task13() {
	a := 5
	b := 23
	fmt.Printf("a = %d b = %d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a = %d b = %d\n", a, b)
}
