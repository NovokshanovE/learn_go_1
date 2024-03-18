package main

import "fmt"

func main() {
	// для того чтобы поменять числа местами, необходимо выполнить следующие действия
	a := 5
	b := 23
	fmt.Printf("a = %d b = %d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a = %d b = %d\n", a, b)
}
