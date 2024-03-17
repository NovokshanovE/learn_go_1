package main

import "fmt"

func solve_t3(a int, res chan int) {

	// fmt.Println(a * a)
	res <- a * a
}

func task3() {
	array := [5]int{2, 4, 6, 8, 10}
	// from := make(chan bool)
	to := make(chan int)
	for _, i := range array {

		go solve_t3(i, to)

	}
	res := 0
	for i := 0; i < 5; i++ {
		res += <-to
	}
	fmt.Println(res)
	// <-from
}
