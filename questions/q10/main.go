package main

import "fmt"

func update(p *int) {
	b := 2
	fmt.Println(p)
	p = &b
	fmt.Println(p)
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
