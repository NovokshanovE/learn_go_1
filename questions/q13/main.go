package main

import "fmt"

func someAction(v []int8, b int8) []int8 {

	v[0] = 100
	v = append(v, b)
	return v
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	a = someAction(a, 6)
	fmt.Println(a)
}
