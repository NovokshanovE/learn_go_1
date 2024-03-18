package main

import "fmt"

func solve_t2(a int, res chan int) { // функция, которая будет отрабатывать в отдельной горутине

	// fmt.Println(a * a)
	res <- a * a
}

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	// from := make(chan bool)
	to := make(chan int) // создание небуферизированного канала
	for _, i := range array {

		go solve_t2(i, to) //запуск горутины

	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-to) //получение данных из канала
	}
	// <-from
}
