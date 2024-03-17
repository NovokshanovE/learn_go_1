package main

import "fmt"

func solve_t2(a int, res chan int) {

	fmt.Println(a * a)
	res <- a * a
}

func task2() {
	array := [5]int{2, 4, 6, 8, 10}
	// from := make(chan bool)
	to := make(chan int)
	for _, i := range array {

		go solve_t2(i, to)

	}
	for i := 0; i < 5; i++ {
		<-to
	}
	// <-from
}
